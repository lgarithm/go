package main

import (
	"flag"
	"log"
	"time"

	"github.com/lgarithm/go/control"
	"github.com/lgarithm/go/net/http/probe"
)

var (
	period = flag.Duration("period", 5*time.Second, "")
)

func main() {
	urls := []string{
		`http://www.qq.com`,
		`http://www.baidu.com`,
		`https://github.com/`,
	}
	var ps []*probe.Probe
	for _, url := range urls {
		p, err := probe.New(url)
		if err != nil {
			log.Printf("invalid URL %s: %v", url, err)
			continue
		}
		ps = append(ps, p)
	}
	control.Periodically(*period, func() {
		for _, p := range ps {
			p.Ping()
		}
	})
}
