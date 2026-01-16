//go:build linux
// +build linux

package cmd

import (
	"fmt"
	"os/exec"
)

func buildCommand(seconds int, task string) *exec.Cmd {
	cmdLine := fmt.Sprintf(
		`sleep %d && %s`,
		seconds, task,
	)

	return exec.Command("bash", "-c", cmdLine)
}
