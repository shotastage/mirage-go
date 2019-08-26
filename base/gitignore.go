package base

import "github.com/shotastage/GFileable"

func mgfileBase() {
	file := GFileable.Path(".gitignore")
	file.WriteString("Mgfile")
	file.WriteString("\n")
}
