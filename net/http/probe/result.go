package probe

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/lgarithm/go/xterm"
)

// SI units
const (
	Ki = 1 << 10
	Mi = 1 << 20
)

var (
	ok   = xterm.Green
	warn = xterm.Yellow
	fail = xterm.Red
)

// Interval represent a time period
type Interval struct {
	Begin time.Time
	End   time.Time
}

// Len returns the length of the Interval
func (i Interval) Len() time.Duration {
	return i.End.Sub(i.Begin)
}

// LonggerThan returns if the duration is logger than given duration
func (i Interval) LonggerThan(t time.Duration) bool {
	return i.Len() > t
}

// Show converts an Interval to human readable string
func (i Interval) Show() string {
	if i.Begin.After(i.End) {
		return "timeout"
	}
	return i.Len().String()
}

// Highlight converts an Interval to human readable string with color thresholds
func (i Interval) Highlight(t1, t2 time.Duration) (string, xterm.Color) {
	switch {
	case i.Begin.After(i.End):
		return "timeout", fail
	case i.LonggerThan(t2):
		return i.Len().String(), fail
	case i.LonggerThan(t1):
		return i.Len().String(), warn
	default:
		return i.Len().String(), ok
	}
}

// Result is a collection of metrics
type Result struct {
	DNSInterval  Interval
	DialInterval Interval
	RWInterval   Interval

	RecvErr   error
	RecvBytes int
}

// RecvRate returns the bytes received in 1s
func (r Result) RecvRate() float64 {
	return float64(r.RecvBytes) / float64(r.RWInterval.Len()) * float64(time.Second)
}

func (r Result) String() string {
	align := func(s string, c xterm.Color) string {
		const fullPad = "                "
		pad := ""
		if l := utf8.RuneCountInString(s); l < len(fullPad) {
			pad = fullPad[:len(fullPad)-l]
		}
		return pad + c.S(s)
	}
	columns := []string{
		fmt.Sprintf("[%s]", r.icron()),
		"dns latency:",
		align(r.DNSInterval.Highlight(5*time.Millisecond, 10*time.Second)),
		"dial latency:",
		align(r.DialInterval.Highlight(10*time.Millisecond, 100*time.Second)),
		"rw latency:",
		align(r.RWInterval.Highlight(100*time.Millisecond, 1*time.Second)),
		"speed:",
		fmt.Sprintf("%-20s", showSpeed(r.RecvRate())),
	}
	return strings.Join(columns, " ")
}

func (r Result) icron() string {
	if r.RecvErr != nil {
		return fail.S("x")
	}
	if r.RecvBytes == 0 {
		return fail.S("x")
	}
	if r.DNSInterval.LonggerThan(1 * time.Second) {
		return warn.S(".")
	}
	if r.RecvRate() < float64(1*Mi) {
		return warn.S(".")
	}
	return xterm.Green.S("v")
}

func showSpeed(rate float64) string {
	return fmt.Sprintf("%f KiB/s", rate/Ki)
}
