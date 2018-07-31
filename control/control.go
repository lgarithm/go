package control

import (
	"context"
	"log"
	"time"
)

func Wait(ctx context.Context, period time.Duration, trial func() error) error {
	if err := trial(); err == nil {
		return nil
	}
	tick := time.Tick(period)
	for {
		select {
		case <-tick:
			if err := trial(); err == nil {
				return nil
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func Try(n int, trial func() error) (err error) {
	for i := 0; i < n; i++ {
		if err = trial(); err == nil {
			return nil
		}
	}
	return
}

func Periodic(d time.Duration, f func()) {
	log.Printf("please use Periodically instead")
	Periodically(d, f)
}

// Periodically run a function periodically at given period
func Periodically(period time.Duration, f func()) {
	f()
	for range time.Tick(period) {
		f()
	}
}
