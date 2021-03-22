package main

import (
	"bytes"
	"fmt"
	"github.com/bukowa/pkgo/src"
	"github.com/bukowa/pkgo/fetcher/git"
	"log"
	"os"
)

func main() {
	src.Registry["git"] = git.Fetcher{}
	b := bytes.NewBuffer([]byte(`
packages:
  - name: test
    type: git
    source: https://github.com/bukowa/pkgo.git
    location: ./test
`))

	c, err := src.NewConfig(b)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	for _, pkg := range c.Packages {
		dir, err := pkg.Fetch(pkg.Package)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(dir)
	}
}
