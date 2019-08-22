package main

import (
	"flag"
	"fmt"
	"mirage-go/info"
	"mirage-go/mgpkg"
	"mirage-go/touch"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("🧐  Too short argument.")
		return
	}

	if args[0] == "bs" || args[0] == "bootstrap" {
		mgpkg.Bootstrap()
		return
	}

	if args[0] == "cr" || args[0] == "create" {
		touch.Touch(args[1])
		return
	}

	if args[0] == "info" {
		info.Process(args[1])
		return
	}

	if args[0] == "v" || args[0] == "version" {
		info.Version()
		return
	}

	if args[0] == "h" || args[0] == "help" {
		info.Help()
		return
	}

	fmt.Println("❌  Given aurguments is invalid.")
}
