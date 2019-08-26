package base

func Process(task string) {

	if task == "ignore-mgfile" {
		mgfileBase()
		return
	}

	println("❌  Invalid argument", task)
}
