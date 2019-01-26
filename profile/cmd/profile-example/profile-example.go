package main

import (
	"os"
	"time"

	"github.com/lgarithm/go/profile"
)

func main() {
	p := profile.New()

	for i := 0; i < 10; i++ {
		func() {
			defer p.Profile("f1").Done()
			time.Sleep(time.Duration(i+1) * time.Millisecond)
		}()
	}
	p.WriteSummary(os.Stdout)
}
