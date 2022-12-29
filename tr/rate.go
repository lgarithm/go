package tr

import (
	"fmt"
	"time"
)

const mi int64 = 1 << 20
const gi int64 = 1 << 30

func rate(w int64, d time.Duration) string {
	if w == 0 {
		return `-`
	}
	r := (float64(w) / float64(gi)) / (float64(d) / float64(time.Second))
	if r < 1 {
		r := (float64(w) / float64(mi)) / (float64(d) / float64(time.Second))
		return fmt.Sprintf("%.3fMiB/s", r)
	}
	return fmt.Sprintf("%.3fGiB/s", r)
}
