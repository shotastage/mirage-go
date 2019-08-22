package info

import (
	"fmt"
	"runtime"

	"github.com/shotastage/GFileable"
)

func Version() {
	fmt.Println()
	fmt.Println("MIRAGE Go Version ", "0.0.1")
	fmt.Println("Copyright (C) 2019 Shota Shimazu All Rights Reserved.")
	fmt.Println("https://lin9.me/WjdSu")
	fmt.Println()
}

func Process(task string) {

	if task == "home" {
		homeDir()
	}

	if task == "os" {
		os()
	}
}

func homeDir() {

	home := GFileable.Path("~").Path

	fmt.Println("🏚  Home directory is ", home)
}

func os() {

	print("⚙️  Your platform is ")

	switch runtime.GOOS {
	case "windows":
		print("Windows")
	case "darwin":
		print("macOS")
	case "linux":
		print("Linux")
	case "freebsd":
		print("FreeBSD")
	default:
		print("Other OS")
	}

	print(".\n")

	return
}
