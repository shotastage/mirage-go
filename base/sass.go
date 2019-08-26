package base

import "github.com/shotastage/GFileable"

func sass() {
	GFileable.Touch("main.scss")

	file := GFileable.Path("main.scss")

	str := `
/**
*
*/
	`
	file.WriteString(str)
}
