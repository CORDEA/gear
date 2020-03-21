package main

import (
	"os"
	"os/exec"
)

func ExecGitCommand(command []string) {
	cmd := exec.Command("git", command...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
