package ios

import "os/exec"

func Bootstrap() {
	println("Bootstrapping iOS... is now under construction")
	out, err := exec.Command("pod", "init").Output()

	if err != nil {
		panic("MIRAGE has been terminated due to unexpected error!")
	}

	println(string(out))
}
