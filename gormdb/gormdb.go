package gormdb

import "github.com/shotastage/GFileable"

func Process(task string) {

	if task == "bs" || task == "bootstrap" {
		bootstrap()
		return
	}

	println("❌  Ivalid argument", task)
}

func bootstrap() {

	if GFileable.Path("models").IsDir() {
		println("Already bootstraped!")
	} else {
		GFileable.Mkdir("models")
	}
}
