package utils

import (
	"os"
	"os/exec"
)

func Clear(osType string) {
	cmd := exec.Command("clear")

	if osType == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
