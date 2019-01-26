package profile

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Duration(f func() error) (time.Duration, error) {
	begin := time.Now()
	err := f()
	d := time.Now().Sub(begin)
	return d, err
}

type profiler struct {
	start time.Time
	name  string
}

func Profile(name string) *profiler {
	fmt.Fprintf(os.Stderr, "deprecated, please use simpleProfiler instead.\n")
	p := profiler{
		name:  name,
		start: time.Now(),
	}
	log.Printf("%s started", p.name)
	return &p
}

func (p *profiler) Done() {
	d := time.Now().Sub(p.start)
	log.Printf("%s took %s", p.name, d)
}
