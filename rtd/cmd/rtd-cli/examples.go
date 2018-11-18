package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/lgarithm/go/rtd"
)

func exampleListProjects(client *rtd.Client) {
	ps, err := client.ListProject(*project)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", ps)
}

func exampleGetProject(client *rtd.Client, id int) {
	project, err := client.GetProject(id)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", project)
}

func exampleListBuilds(client *rtd.Client) {
	builds, err := client.ListBuild(*project)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	for _, b := range builds {
		showBuild(&b)
	}
}

func exampleGetBuild(client *rtd.Client, id int) {
	build, err := client.GetBuild(id)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	showBuild(build)
}

func exampleGetBuildLog(client *rtd.Client, id int) {
	rawLog, err := client.GetBuild(id)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	fmt.Print(rawLog)
}

func exampleListVersions(client *rtd.Client) {
	versions, err := client.ListVersion(*project)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	for _, v := range versions {
		showVersion(&v)
	}
}

func exampleGetVersion(client *rtd.Client, id int) {
	version, err := client.GetVersion(id)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	showVersion(version)
}

func showBuild(b *rtd.Build) {
	fmt.Printf("#%d\n", b.ID)
	for _, c := range b.Commands {
		fmt.Printf("    %s\n", c.Command)
	}
}

func showVersion(v *rtd.Version) {
	fmt.Printf("#%d\n", v.ID)
	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "    ")
	e.Encode(v)
}
