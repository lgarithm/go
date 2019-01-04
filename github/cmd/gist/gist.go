package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/lgarithm/go/github"
)

var (
	name       = flag.String("n", "new gist", "name of the gist")
	file       = flag.String("f", "", "path to file")
	content    = flag.String("c", "", "content as string")
	listRecent = flag.Int("l", 0, "list most recent l gists")
	downloadID = flag.String("d", "", "ID to download")
)

func main() {
	flag.Parse()
	if *listRecent > 0 {
		if err := listGist(*listRecent); err != nil {
			exitErr(err)
		}
		return
	}
	if len(*downloadID) > 0 {
		exitErr(errors.New("Not implemented"))
		return
	}
	files := make(map[string]string)
	if len(*file) > 0 {
		bs, err := ioutil.ReadFile(*file)
		if err == nil {
			files[path.Base(*file)] = string(bs)
		}
	}
	if len(*content) > 0 {
		files[*name] = *content
	}
	if len(files) <= 0 {
		fmt.Printf("No Files nor Content!\n")
		flag.Usage()
		os.Exit(1)
	}
	if err := createGist(*name, files); err != nil {
		exitErr(err)
	}
}

func listGist(n int) error {
	githubClient, err := github.NewSimpleClient()
	if err != nil {
		return err
	}
	gs, err := githubClient.ListGist()
	if err != nil {
		return err
	}
	if len(gs) > n {
		gs = gs[:n]
	}
	now := time.Now()
	for _, g := range gs {
		fmt.Printf("%s %s (updated %s ago)\n", g.ID, g.Description, now.Sub(g.UpdatedAt))
		for name := range g.Files {
			fmt.Printf("    %s\n", name)
		}
	}
	return nil
}

func createGist(name string, files map[string]string) error {
	githubClient, err := github.NewSimpleClient()
	if err != nil {
		return err
	}

	if err := githubClient.CreateGist(name, files); err != nil {
		return err
	}
	return nil
}

func exitErr(err error) {
	log.Print(err)
	os.Exit(1)
}
