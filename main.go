package main

import (
	"flag"
	"mirage-go/mgpkg"
	"mirage-go/touch"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if args[0] == "bs" || args[0] == "bootstrap" {
		mgpkg.Bootstrap()
	}

	if args[0] == "cr" || args[0] == "create" {
		touch.Touch(args[1])
	}
}
