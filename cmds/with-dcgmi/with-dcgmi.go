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
	"time"

	"github.com/lgarithm/go/monitoring"
)

var (
	period = flag.Duration(`p`, 100*time.Millisecond, `monitoring period`)
	logDir = flag.String(`logdir`, `logs`, `log dir`)
)

func main() {
	flag.Parse()
	m := monitoring.NewMonitor(*period)
	m.LogDir = *logDir
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
