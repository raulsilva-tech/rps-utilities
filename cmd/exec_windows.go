//go:build windows
// +build windows

package cmd

import (
	"fmt"
	"os/exec"
	"syscall"
)

func buildCommand(seconds int, task string) *exec.Cmd {
	cmdLine := fmt.Sprintf(
		`timeout /T %d /NOBREAK >NUL && %s`,
		seconds, task,
	)

	cmd := exec.Command("cmd.exe", "/C", cmdLine)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: CREATE_NEW_PROCESS_GROUP | DETACHED_PROCESS,
	}

	return cmd
}
