package tr

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

var UseXterm = true

type patient struct {
	threshold time.Duration
	done      int32
	slow      int32
	name      string
	t0        time.Time
	xterm     bool
}

func (p *patient) Done() {
	atomic.StoreInt32(&p.done, 1)
	if atomic.LoadInt32(&p.slow) > 0 {
		ok := `finially finished`
		if p.xterm {
			ok = Green.S(ok)
		}
		fmt.Fprintf(os.Stderr, "%s %s, took %s\n", ok, p.name, time.Since(p.t0))
	}
}

func (p *patient) start() {
	warn := `still running`
	if p.xterm {
		warn = Yellow.S(warn)
	}
	for {
		if atomic.LoadInt32(&p.done) > 0 {
			break
		}
		time.Sleep(p.threshold)
		atomic.StoreInt32(&p.slow, 1)
		if atomic.LoadInt32(&p.done) == 0 {
			fmt.Fprintf(os.Stderr, "%s %s, took %s\n", warn, p.name, time.Since(p.t0))
		}
	}
}

func Patient(name string, d time.Duration) *patient {
	p := &patient{
		name:      name,
		t0:        time.Now(),
		threshold: d,
		xterm:     UseXterm,
	}
	go p.start()
	return p
}
