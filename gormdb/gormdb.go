package gormdb

import "github.com/shotastage/GFileable"

func Bootstrap() {

	if GFileable.Path("models").IsDir() {
		println("Already bootstraped!")
	} else {
		GFileable.Mkdir("models")
	}
}
