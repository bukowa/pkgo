package main

import (
	"bytes"
	"fmt"
	"github.com/bukowa/pkgo/src"
	"github.com/bukowa/pkgo/types/git"
	"os"
)

func main() {
	src.Registry["git"] = git.Fetcher{}
	b := bytes.NewBuffer([]byte(`
packages:
  - name: test
    type: git`))

	c, err := src.NewConfig(b)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	if x := c.Packages[0]; x.Type == "git" {
		s, err := x.Fetch(x.Package)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		if s == "git" {
			os.Exit(0)
		}
	}
	panic("")
}

