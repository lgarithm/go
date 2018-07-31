package main

import (
	"flag"
	"time"

	"github.com/lgarithm/go/control"
	"github.com/lgarithm/go/net/http/probe"
)

var (
	period = flag.Duration("period", 1*time.Second, "")
)

func main() {
	urls := []string{
		`http://www.qq.com`,
		`http://www.baidu.com`,
	}
	var ps []*probe.Probe
	for _, url := range urls {
		p := probe.New(url)
		ps = append(ps, p)
	}
	control.Periodically(*period, func() {
		for _, p := range ps {
			p.Ping()
		}
	})
}
