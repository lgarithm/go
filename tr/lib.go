package tr

import (
	"flag"
	"fmt"
	"io"
	"strings"
	"time"
)

var reportFile = flag.String("trace-file", "profile.log", "")

type Scope interface{ Done() }

var defaultCtx Ctx = &simpleCtx{
	baseCtx: baseCtx{
		t0: time.Now(),
		cs: newCounters(),
	},
}

func Set(c Ctx) { defaultCtx = c }

func SetTitle(t string) { defaultCtx.SetTitle(t) }

func Reset() { defaultCtx.Reset() }

func Report(w io.Writer) { defaultCtx.Report(w) }

func ReportFile(filename string) { defaultCtx.ReportFile(filename) }

func Save() { defaultCtx.ReportFile(*reportFile) }

func Trace(name string) Scope { return defaultCtx.Trace(name) }

func TraceW(name string, w int64) Scope { return defaultCtx.TraceW(name, w) }

func TraceF(name string, f func()) {
	defer Trace(name).Done()
	f()
}

func TraceFnErr(name string, f func() error) error {
	defer Trace(name).Done()
	return f()
}

type showD time.Duration

func (d showD) String() string {
	switch dd := time.Duration(d); {
	case dd < time.Microsecond:
		return fmt.Sprintf("%dns", dd/time.Nanosecond)
	case dd < time.Millisecond:
		return fmt.Sprintf("%7.3fus", float64(dd)/float64(time.Microsecond))
	case dd < time.Second:
		return fmt.Sprintf("%7.3fms", float64(dd)/float64(time.Millisecond))
	case dd < time.Minute:
		return fmt.Sprintf("%7.3fs", float64(dd)/float64(time.Second))
	default:
		return dd.String()
	}
}

var hr = strings.Repeat(`-`, 80)
