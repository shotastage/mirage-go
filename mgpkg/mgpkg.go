package mgpkg

import (
	"mirage-go/cmd"

	"github.com/shotastage/GFileable"
)

func Bootstrap() {
	println("Initializing MIRAGE package...")
	GFileable.Touch("Mgfile")

	file := GFileable.Path("./Mgfile")
	file.WriteString(":* This file is generated by Mirage Go *:\n")
	file.WriteString("name={{PROJECT_NAME}}\n")
	file.WriteString("author={{AUTHOR_NAME}}\n")
	file.WriteString("version={{PROJECT_VERSION}}\n")
	file.WriteString("language={{PROJECT_LANGUAGE}}\n")

	cmd.Sess.Exec()
}
