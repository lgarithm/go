package profile

import (
	"fmt"
	"io"
	"sort"
	"sync"
	"time"
)

type simpleScope struct {
	Scope

	name     string
	start    time.Time
	profiler *simpleProfiler
}

type simpleProfiler struct {
	sync.Mutex

	counts         map[string]int64
	minDurations   map[string]time.Duration
	maxDurations   map[string]time.Duration
	totalDurations map[string]time.Duration
}

func (s *simpleScope) Done() {
	s.profiler.done(s, time.Now())
}

func NewSimpleProfiler() Profiler {
	return &simpleProfiler{
		counts:         make(map[string]int64),
		minDurations:   make(map[string]time.Duration),
		maxDurations:   make(map[string]time.Duration),
		totalDurations: make(map[string]time.Duration),
	}
}

func (p *simpleProfiler) Profile(name string) Scope {
	return &simpleScope{
		name:     name,
		start:    time.Now(),
		profiler: p,
	}
}

func (p *simpleProfiler) done(s *simpleScope, t time.Time) {
	p.Lock()
	defer p.Unlock()
	d := t.Sub(s.start)

	p.counts[s.name]++
	p.totalDurations[s.name] += d
	if val, ok := p.minDurations[s.name]; !ok || d < val {
		p.minDurations[s.name] = d
	}
	if val, ok := p.maxDurations[s.name]; !ok || d > val {
		p.maxDurations[s.name] = d
	}
}

func (p *simpleProfiler) WriteSummary(w io.Writer) error {
	p.Lock()
	defer p.Unlock()

	var names []string
	for name := range p.counts {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool { return p.totalDurations[names[i]] > p.totalDurations[names[j]] })

	th := []string{"count", "mean", "min", "max", "total", "call site"}
	var rows [][]string
	for _, name := range names {
		rows = append(rows, []string{
			fmt.Sprintf("%d", p.counts[name]),
			fmt.Sprintf("%s", p.totalDurations[name]/time.Duration(p.counts[name])),
			fmt.Sprintf("%s", p.minDurations[name]),
			fmt.Sprintf("%s", p.maxDurations[name]),
			fmt.Sprintf("%s", p.totalDurations[name]),
			fmt.Sprintf("%s", name),
		})
	}
	return showTable(w, th, rows)
}
