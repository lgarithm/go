package probe

import (
	"bytes"
	"fmt"
	"time"
)

// SI units
const (
	Ki = 1 << 10
)

// Interval represent a time period
type Interval struct {
	Begin time.Time
	End   time.Time
}

// Len return the length of the Interval
func (i Interval) Len() time.Duration {
	return i.End.Sub(i.Begin)
}

// Show converts an Interval to human readable string
func (i Interval) Show() string {
	if i.Begin.After(i.End) {
		return "timeout"
	}
	return i.Len().String()
}

// Result is a collection of metrics
type Result struct {
	DNSInterval  Interval
	DialInterval Interval
	RWInterval   Interval

	RecvErr   error
	RecvBytes int
}

func (r Result) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "dns latency: %s", r.DNSInterval.Show())
	fmt.Fprint(&b, ", ")
	fmt.Fprintf(&b, "dial latency: %s", r.DialInterval.Show())
	fmt.Fprint(&b, ", ")

	rate := float64(r.RecvBytes) / float64(r.RWInterval.Len()) * float64(time.Second)
	fmt.Fprintf(&b, "rw latency: %s, speed: %f KiB/s", r.RWInterval.Len(), rate/Ki)
	return b.String()
}
