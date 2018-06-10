package supervisord

import (
	"context"
	"io"
	"log"
	"os/exec"

	"github.com/lgarithm/go/iostream"
)

// Supervisord runs an external program
type Supervisord struct {
	name    string
	args    []string
	ctx     context.Context
	cancel  context.CancelFunc
	stopped chan struct{}
}

// New creates a new Supervisord
func New(name string, args ...string) *Supervisord {
	log.Printf("New Supervisord for %s with %q", name, args)
	ctx, cancel := context.WithCancel(context.TODO())
	return &Supervisord{
		name:    name,
		args:    args,
		ctx:     ctx,
		cancel:  cancel,
		stopped: make(chan struct{}, 1),
	}
}

// SupervisedRun keeps the programming running
func (s *Supervisord) SupervisedRun() {
	for {
		select {
		case <-s.ctx.Done():
			s.stopped <- struct{}{}
			return
		default:
		}
		if err := s.run(s.ctx); err != nil {
			if err == context.Canceled {
				s.stopped <- struct{}{}
				return
			}
			log.Printf("%s failed with: %v", s.name, err)
		}
	}
}

// Stop cancels the program
func (s *Supervisord) Stop() {
	s.cancel()
	<-s.stopped
}

func (s *Supervisord) run(ctx context.Context) error {
	cmd := exec.Command(s.name, s.args...)
	if stdout, err := cmd.StdoutPipe(); err == nil {
		go streamPipe("stdout", stdout)
	} else {
		return err
	}
	if stderr, err := cmd.StderrPipe(); err == nil {
		go streamPipe("stderr", stderr)
	} else {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	done := make(chan error)
	go func() { done <- cmd.Wait() }()
	select {
	case <-ctx.Done():
		cmd.Process.Kill()
		return cmd.Wait()
		// return ctx.Err()
	case err := <-done:
		return err
	}
}

func streamPipe(name string, r io.Reader) error {
	w := iostream.NewLogWriter(name)
	return iostream.StreamPipe(r, w)
}
