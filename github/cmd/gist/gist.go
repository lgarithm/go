package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/lgarithm/go/github"
)

var (
	name    = flag.String("n", "new gist", "name of the gist")
	file    = flag.String("f", "", "path to file")
	content = flag.String("c", "", "content as string")
)

func main() {
	flag.Parse()
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
		log.Print(err)
		os.Exit(1)
	}
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
