package main

import (
	"fmt"

	"github.com/lgarithm/go/git"
)

func main() {
	h := git.HashObject([]byte("# README\n"))
	fmt.Printf("%s\n", h)
}
