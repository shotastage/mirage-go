package ios

import "os/exec"

func command(cmd string) error {
	out, err := exec.Command("ls", "-la").Output()
	println(out)

	return err
}

func Bootstrap() {
	println("Bootstrapping iOS... is now under construction")
	command("pod init")
}
