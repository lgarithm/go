/*
Usage:

make -C benchmark
./benchmark/bin/with-dcgmi <cmd> [args]
*/
package main

import (
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/lgarithm/go/monitoring"
)

func main() {
	flag.Parse()

	m := monitoring.NewMonitor()
	m.Start()
	defer m.Stop()

	args := flag.Args()
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("running %s", cmd)
	if err := cmd.Run(); err != nil {
		log.Panicln(err)
	}
}
