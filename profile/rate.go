package profile

import (
	"fmt"
	"io"
	"sort"
	"sync"
	"time"
)

type rateScope struct {
	name     string
	start    time.Time
	amount   int64
	profiler *rateProfiler
}

type rateProfiler struct {
	sync.Mutex

	counts         map[string]int64
	minRates       map[string]float64
	maxRates       map[string]float64
	totalAmounts   map[string]int64
	totalDurations map[string]time.Duration
}

func (s *rateScope) Done() {
	s.profiler.done(s, time.Now())
}

func (s *rateScope) WithAmount(n int64) Scope {
	s.amount = n
	return s
}

func NewRateProfiler() Profiler {
	return &rateProfiler{
		counts:         make(map[string]int64),
		minRates:       make(map[string]float64),
		maxRates:       make(map[string]float64),
		totalAmounts:   make(map[string]int64),
		totalDurations: make(map[string]time.Duration),
	}
}

func (p *rateProfiler) Profile(name string) Scope {
	return &rateScope{
		name:     name,
		start:    time.Now(),
		profiler: p,
	}
}

func rate(a int64, d time.Duration) float64 {
	return float64(a) * float64(time.Second) / float64(d)
}

func showRate(r float64) string {
	switch {
	case r < 1e3:
		return fmt.Sprintf("%.2f/s", r)
	case r < 1e6:
		return fmt.Sprintf("%.2fk/s", r/1e3)
	default:
		return fmt.Sprintf("%.2fm/s", r/1e6)
	}
}

func (p *rateProfiler) done(s *rateScope, t time.Time) {
	p.Lock()
	defer p.Unlock()
	d := t.Sub(s.start)
	r := rate(s.amount, d)

	p.counts[s.name]++
	p.totalAmounts[s.name] += s.amount
	p.totalDurations[s.name] += d
	if val, ok := p.minRates[s.name]; !ok || r < val {
		p.minRates[s.name] = r
	}
	if val, ok := p.maxRates[s.name]; !ok || r > val {
		p.maxRates[s.name] = r
	}
}

func (p *rateProfiler) WriteSummary(w io.Writer) error {
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
			fmt.Sprintf("%s", showRate(rate(p.totalAmounts[name], p.totalDurations[name]))),
			fmt.Sprintf("%s", showRate(p.minRates[name])),
			fmt.Sprintf("%s", showRate(p.maxRates[name])),
			fmt.Sprintf("%s", p.totalDurations[name]),
			fmt.Sprintf("%s", name),
		})
	}
	return showTable(w, th, rows)
}
