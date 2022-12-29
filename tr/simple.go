package tr

import (
	"fmt"
	"io"
	"os"
	"sort"
	"sync"
	"time"
)

type simpleCtx struct {
	baseCtx

	mu sync.RWMutex
}

func Simple() *simpleCtx {
	return &simpleCtx{
		baseCtx: baseCtx{
			t0: time.Now(),
			cs: newCounters(),
		},
	}
}

func (c *simpleCtx) In(name string) {}

func (c *simpleCtx) Out(name string, d time.Duration, w int64) {
	c.mu.Lock()
	c.cs.add(name, d, w)
	c.mu.Unlock()
}

func (c *simpleCtx) Trace(name string) Scope {
	return c.TraceW(name, 0)
}

func (c *simpleCtx) TraceW(name string, w int64) Scope {
	s := &scope{ctx: c, t0: time.Now(), name: name, w: w}
	c.In(name)
	return s
}

func (s *simpleCtx) Report(w io.Writer) {
	s.report(w, time.Since(s.t0))
}

func (s *simpleCtx) ReportFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	s.Report(f)
	return nil
}

func (s *simpleCtx) Reset() {
	s.t0 = time.Now()
	s.cs = newCounters()
}

func (c *simpleCtx) report(w io.Writer, d time.Duration) {
	var names []string
	for name := range c.cs.cs {
		names = append(names, name)
	}

	sort.Slice(names, func(i, j int) bool {
		a := c.cs.cs[names[i]]
		b := c.cs.cs[names[j]]
		return a.d > b.d
	})

	title(w, c.title)
	fmt.Fprintf(w, "%12s %16s %16s  %8s %16s | %s\n", `#`, `tot`, `mean`, `%`, `~    `, `name`)
	fmt.Fprintf(w, "%s\n", hr)
	for _, name := range names {
		c := c.cs.cs[name]
		fmt.Fprintf(w, "%12d %16s %16s %8.3f%% %16s | %s\n", c.n, showD(c.d), showD(c.d/time.Duration(c.n)), 100.0*float64(c.d)/float64(d), rate(c.w, c.d), name)
	}
}
