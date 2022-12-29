package tr

import (
	"io"
	"time"
)

type Ctx interface {
	In(name string)
	Out(name string, d time.Duration, w int64)
	Trace(name string) Scope
	TraceW(name string, w int64) Scope

	SetTitle(t string)
	Reset()
	Report(io.Writer)
	ReportFile(string) error
}

type baseCtx struct {
	title string
	t0    time.Time
	cs    *counters
}

func (c *baseCtx) SetTitle(t string) { c.title = t }
