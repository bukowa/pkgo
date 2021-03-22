package main

import (
	"fmt"
	v1git "github.com/bukowa/pkgo/fetcher/git/v1"
	"github.com/bukowa/pkgo/fetcher/wporg/v1"
	"github.com/bukowa/pkgo/src"
	"log"
	"os"
)

func main() {
	f, err := os.Open("example.yaml")
	if err != nil {
		panic(err)
	}

	src.Registry["git.v1"] = v1git.Fetcher{}
	src.Registry["wp_theme.v1"] = v1wporg.Fetcher{Kind: v1wporg.KindTheme}
	src.Registry["wp_plugin.v1"] = v1wporg.Fetcher{Kind: v1wporg.KindPlugin}

	c, err := src.NewConfig(f)
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
