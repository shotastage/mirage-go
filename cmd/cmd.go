package cmd

import (
	"os/exec"
)

func isCommandAvailable(name string) bool {
	cmd := exec.Command(name)

	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}
