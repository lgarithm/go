package tr

import (
	"time"
)

type counter struct {
	n int
	d time.Duration
	w int64
}

type counters struct {
	cs map[string]*counter
}

func newCounters() *counters {
	return &counters{cs: make(map[string]*counter)}
}

func (c *counters) add(name string, d time.Duration, w int64) {
	p, ok := c.cs[name]
	if !ok {
		p = new(counter)
		c.cs[name] = p
	}

	p.n += 1
	p.d += d
	p.w += w
}
