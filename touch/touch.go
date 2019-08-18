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
		println("Creating Python file")

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
		println("Creating Go file")
	}

	if strings.Contains(ext[1], "sh") {
		println("Creating Shell file")
	}

	if strings.Contains(ext[1], "js") {
		println("Creating JavaScript file")
	}
}
