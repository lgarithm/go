package control

import (
	"context"
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
	f()
	for _ = range time.Tick(d) {
		f()
	}
}
