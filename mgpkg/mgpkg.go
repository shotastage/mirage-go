package mgpkg

import (
	"github.com/shotastage/GFileable"
	"mirage-go/cmd"
)

func Bootstrap() {
	println("Initializing MIRAGE package...")
	GFileable.Touch("Mgfile")

	file := GFileable.Path("")
	file.WriteString("DD")

	cmd.Sess.Exec()
}
