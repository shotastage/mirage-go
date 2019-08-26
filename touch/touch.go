package touch

import (
	"strings"

	"github.com/shotastage/GFileable"
)

func Touch(file string) {

	ext := strings.Split(file, ".")

	if len(ext) < 2 {
		println("Please type following format...")
		return
	}

	if strings.Contains(ext[1], "py") {

		GFileable.Touch(file)

		file := GFileable.Path(file)
		file.WriteString("# FILE_NAME \n")
		file.WriteString("# Created by AUTHOR_NAME on DATE\n")
		file.WriteString("# Copyright © 2019 COPYRIGHT_NAME\n")
		file.WriteString("#\n")
		file.WriteString("\n")
		file.WriteString("import os\n")
		file.WriteString("\n")
		file.WriteString("print(\"Hello, world!\")\n")
		file.WriteString("\n")
	}

	if strings.Contains(ext[1], "go") {
		GFileable.Touch(file)

		file := GFileable.Path(file)
		file.WriteString("// FILE_NAME \n")
		file.WriteString("// Created by AUTHOR_NAME on DATE\n")
		file.WriteString("// Copyright © 2019 COPYRIGHT_NAME\n")
		file.WriteString("//\n")
		file.WriteString("\n")
		file.WriteString("package pkg_name\n")
		file.WriteString("\n")
	}

	if strings.Contains(ext[1], "sh") {
		println("🙇  Sorry creating Shell file does not supported now.")
	}

	if strings.Contains(ext[1], "js") {
		println("🙇  Sorry creating JavaScript file does not supported now.")
	}

	if strings.Contains(ext[1], "swift") {
		println("🙇  Sorry creating Swift file does not supported now.")
	}
}
