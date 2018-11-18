package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/lgarithm/go/control"
	"github.com/lgarithm/go/net/http/probe"
)

var (
	defaultConfig = path.Join(os.Getenv("HOME"), ".http-ping")

	period  = flag.Duration("period", 1*time.Second, "")
	timeout = flag.Duration("timeout", 5*time.Second, "")
	config  = flag.String("config", defaultConfig, "path to config file")
)

func main() {
	flag.Parse()
	urls := loadConfig(*config)
	var ps []*probe.Probe
	for _, url := range urls {
		p, err := probe.New(url, *timeout)
		if err != nil {
			log.Printf("invalid URL %s: %v", url, err)
			continue
		}
		ps = append(ps, p)
	}
	control.Periodically(time.Duration(len(ps))*(*period), func() {
		for _, p := range ps {
			p.Ping()
		}
	})
}

var defaultURLs = []string{
	`http://www.qq.com`,
	`http://www.baidu.com`,
}

func loadConfig(filename string) []string {
	var urls []string
	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Config %s not found, Using default config", filename)
		return defaultURLs
	}
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			break
		}
		{
			line := strings.TrimSpace(string(line))
			if len(line) == 0 || strings.HasPrefix(line, "#") {
				continue
			}
			u, err := url.Parse(line)
			if err != nil {
				log.Printf("invalid URL in config: %s, ignored", line)
				continue
			}
			urls = append(urls, u.String())
		}
	}
	return urls
}
