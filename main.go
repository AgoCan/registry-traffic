package main

import (
	"flag"
	"os"
)

var (
	isShow = flag.Bool("v", false, "show version")
)

func main() {
	flag.Parse()
	if *isShow {
		showVersion()
		os.Exit(0)
	}
}
