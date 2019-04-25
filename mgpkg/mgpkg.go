package mgpkg

import "github.com/shotastage/GFileable"

func BootstrapMiragePackage() {
	print("Initializing MIRAGE package...")
	GFileable.Touch("Mgpkg")
}
