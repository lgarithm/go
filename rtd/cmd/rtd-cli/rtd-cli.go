package main

import (
	"flag"

	"github.com/lgarithm/go/rtd"
)

var (
	project = flag.String("project", `tensorlayer`, "project name")
)

func main() {
	flag.Parse()
	client := rtd.New()

	client.Ping()

	// exampleListProjects(client)
	// const projectID = 56541
	// exampleGetProject(client, projectID)

	// exampleListBuilds(client)
	// const buildID = 7577722
	// exampleGetBuild(client, buildID)
	// exampleGetBuildLog(client, buildID)

	exampleListVersions(client)
	// exampleGetVersion(client, 0)
}
