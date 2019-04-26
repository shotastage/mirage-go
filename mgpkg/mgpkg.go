package mgpkg

import (
	"github.com/shotastage/GFileable"
	"mirage-go/cmd"
)

func BootstrapMiragePackage() {
	print("Initializing MIRAGE package...")
	GFileable.Touch("Mgpkg")

	cmd.Sess.Exec()
}
