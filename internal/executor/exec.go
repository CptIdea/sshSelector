package executor

import (
	"os"
	"os/exec"
	"strings"
)

func Connect(host string) error {
	cmd := exec.Command("ssh", strings.Split(host," ")...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
