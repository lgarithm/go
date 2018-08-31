package main

import (
	"flag"

	"github.com/lgarithm/go/travisci/v3"
)

func main() {
	flag.Parse()
	client := travisci.New()
	client.Ping()
}
