package info

import (
	"fmt"

	"github.com/shotastage/GFileable"
)

func Process(task string) {

	if task == "home" {
		homeDir()
	}
}

func homeDir() {

	home := GFileable.Path("~").Path

	fmt.Println("🏚  Home directory is ", home)
}
