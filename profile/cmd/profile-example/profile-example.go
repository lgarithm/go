package main

import (
	"os"
	"time"

	"github.com/lgarithm/go/profile"
)

func main() {
	example1()
	example2()
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
