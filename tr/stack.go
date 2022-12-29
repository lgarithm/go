package tr

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type stackCtx struct {
	baseCtx

	ns map[string]int
	sn []string

	p []int
}

func Stack() *stackCtx {
	return &stackCtx{
		baseCtx: baseCtx{
			t0: time.Now(),
			cs: newCounters(),
		},
		ns: make(map[string]int),
	}
}

func (c *stackCtx) id(name string) int {
	if i, ok := c.ns[name]; ok {
		return i
	}
	i := len(c.ns)
	c.ns[name] = i
	c.sn = append(c.sn, name)
	return i
}

func (c *stackCtx) decode(p []int) string {
	var ns []string
	for _, i := range p {
		ns = append(ns, c.sn[i])
	}
	return strings.Join(ns, `/`)
}

func (c *stackCtx) In(name string) {
	c.p = append(c.p, c.id(name))
}

func (c *stackCtx) Out(name string, d time.Duration, w int64) {
	k := ints(c.p)
	c.cs.add(k, d, w)
	c.p = c.p[:len(c.p)-1]
}

func (s *stackCtx) Reset() {
	s.t0 = time.Now()
	s.cs = newCounters()
	s.ns = make(map[string]int)
	s.sn = nil
}

func (c *stackCtx) Trace(name string) Scope {
	return c.TraceW(name, 0)
}

func (c *stackCtx) TraceW(name string, w int64) Scope {
	s := &scope{ctx: c, t0: time.Now(), name: name, w: w}
	c.In(name)
	return s
}

func (c *stackCtx) sortKeys() []string {
	var names []string
	for name := range c.cs.cs {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		pa := parse(names[i])
		pb := parse(names[j])
		for k := 0; k < len(pa) && k < len(pb); k++ {
			a := c.cs.cs[ints(pa[:k+1])]
			b := c.cs.cs[ints(pb[:k+1])]
			if a == nil || b == nil { // only happens when not finshed
				if a == nil && b == nil {
					continue
				}
				return true
			}
			if a.d > b.d {
				return true
			}
			if a.d < b.d {
				return false
			}
		}
		a := c.cs.cs[names[i]]
		b := c.cs.cs[names[j]]
		if a == nil || b == nil { // only happens when not finshed
			return true
		}
		return a.d > b.d
	})
	return names
}

func (c *stackCtx) Report(w io.Writer) {
	d := time.Since(c.t0)
	names := c.sortKeys()
	title(w, c.title)
	fmt.Fprintf(w, "%12s %16s %16s  %8s %16s | %s\n", `#`, `tot`, `mean`, `%`, `~    `, `name`)
	fmt.Fprintf(w, "%s\n", hr)
	for _, name := range names {
		p := c.decode(parse(name))
		c := c.cs.cs[name]
		fmt.Fprintf(w, "%12d %16s %16s %8.3f%% %16s | %s\n", c.n, showD(c.d), showD(c.d/time.Duration(c.n)), 100.0*float64(c.d)/float64(d), rate(c.w, c.d), p)
	}
}

func (s *stackCtx) ReportFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	s.Report(f)
	return nil
}

func ints(xs []int) string {
	w := &bytes.Buffer{}
	for i, x := range xs {
		if i > 0 {
			fmt.Fprintf(w, `,`)
		}
		fmt.Fprintf(w, "%d", x)
	}
	return w.String()
}

func parse(l string) []int {
	var xs []int
	for _, s := range strings.Split(l, `,`) {
		x, err := strconv.Atoi(s)
		if err == nil {
			xs = append(xs, x)
		}
	}
	return xs
}
