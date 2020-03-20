package main

import (
	"os"
	"os/exec"
)

type Add struct {
}

func (a *Add) equals(source string) bool {
	return source == "a" || source == "add"
}

func (a *Add) exec(args []string) {
	cmd := exec.Command("git", append([]string{"add"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
