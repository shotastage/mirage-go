package info

import (
	"fmt"
	"mirage-go/shared"
	"runtime"

	"github.com/shotastage/GFileable"
)

func Version() {
	fmt.Println("MIRAGE Go Version ", "0.0.1")
	fmt.Println("Copyright (C) 2019 Shota Shimazu All Rights Reserved.")
	fmt.Println("https://lin9.me/WjdSu")
	fmt.Println()
}

func Process(task string) {

	if task == "help" {
		infoHelp()
		return
	}

	if task == "home" {
		homeDir()
		return
	}

	if task == "os" {
		os()
		return
	}

	if task == "config-file-path" {
		fab()
		return
	}

	if task == "check-compatibility" {
		compatibility()
		return
	}

	println("❌  Ivalid argument", task)
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

func fab() {
	print("🛠  MIRAGE User configuration path is　")
	file := GFileable.Join(shared.UserConfigPath, "UserConfig.json")

	print(file.Path, "\n")
}
