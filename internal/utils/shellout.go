package utils

import (
	"os"
	"os/exec"
)

const ShellToUse string = "bash"

func Shellout(command string) error {
	cmd := exec.Command(ShellToUse, "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	return err
}
