package cmd

import (
	"os/exec"
)

// CommandObject is Command Session struct
type CommandObject struct {
	command string
	output  string
	result  bool
}

// Exec is function that execute commands
func (c *CommandObject) Exec() bool {
	cmd := exec.Command("", "", "")

	if err := cmd.Run(); err != nil {
		return false
	}

	defer func() {
		println("Closing session...")
	}()

	return true
}

// SetCommand is set command
func (c *CommandObject) SetCommand(cmd string) {
	c.command = cmd
}

var Sess *CommandObject
