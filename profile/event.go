package profile

import (
	"fmt"
	"io"
	"sync"
	"time"
)

type event struct {
	name    string
	startNs int64
	endNs   int64
}

type eventScope struct {
	Scope

	name     string
	start    time.Time
	profiler *eventProfiler
}

func (s *eventScope) Done() {
	s.profiler.done(s, time.Now())
}

type eventProfiler struct {
	sync.Mutex

	events []event
}

func NewEventProfiler() Profiler {
	return &eventProfiler{}
}

func (p *eventProfiler) Profile(name string) Scope {
	return &eventScope{
		name:     name,
		start:    time.Now(),
		profiler: p,
	}
}

func (p *eventProfiler) done(s *eventScope, t time.Time) {
	p.Lock()
	defer p.Unlock()
	p.events = append(p.events, event{
		name:    s.name,
		startNs: s.start.UnixNano(),
		endNs:   t.UnixNano(),
	})
}

func (p *eventProfiler) WriteSummary(w io.Writer) error {
	p.Lock()
	defer p.Unlock()
	for _, e := range p.events {
		fmt.Fprintf(w, "%d %d %s\n", e.startNs, e.endNs, e.name)
	}
	return nil
}
