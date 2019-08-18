package main

import (
	"flag"
	"fmt"
	"mirage-go/mgpkg"
	"mirage-go/os"
	"mirage-go/touch"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if args[0] == "bs" || args[0] == "bootstrap" {
		mgpkg.Bootstrap()
		return
	}

	if args[0] == "cr" || args[0] == "create" {
		touch.Touch(args[1])
		return
	}

	if args[0] == "chk-pltfm" || args[0] == "check-platform" {
		os.CheckPlatform()
		return
	}

	fmt.Println("Given aurguments is invalid.")
}
