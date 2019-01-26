package main

import (
	"os"
	"time"

	"github.com/lgarithm/go/profile"
)

func main() {
	example1()
	example2()
	example3()
}

func example1() {
	p := profile.New()
	defer p.WriteSummary(os.Stdout)
	for i := 0; i < 10; i++ {
		func() {
			defer p.Profile("f1").Done()
			time.Sleep(time.Duration(i+1) * time.Millisecond)
		}()
	}
}

func example2() {
	p := profile.NewEventProfiler()
	defer p.WriteSummary(os.Stdout)
	for i := 0; i < 10; i++ {
		func() {
			defer p.Profile("f1").Done()
			time.Sleep(time.Duration(i+1) * time.Millisecond)
		}()
	}
}

func example3() {
	const Mi = 1 << 20
	p := profile.NewRateProfiler()
	defer p.WriteSummary(os.Stdout)
	for i := 10; i < 100; i += 5 {
		func() {
			defer p.Profile("f1").WithAmount(int64(i * Mi)).Done()
			time.Sleep(time.Duration(i) * time.Millisecond)
		}()
	}
}
