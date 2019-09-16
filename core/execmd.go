package core

import (
	"fmt"
	"os/exec"
)

func isCommandAvailable(name string) bool {
	cmd := exec.Command(name, "--help")
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
