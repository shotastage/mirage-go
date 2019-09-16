package runner

import (
	"fmt"
	"os/exec"
)

func Process() {
	println("🏃  Running project...")
	println("🔄  Building project...")
	err := exec.Command("go", "build").Run()
	if err != nil {
		fmt.Println("❌  Failed to build project!")
	}

	run := exec.Command("./bhaa-cloud").Run()
	if run != nil {
		fmt.Println("❌  Failed to start project!")
	}
}
