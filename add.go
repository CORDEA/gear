package main

import (
	"os"
	"os/exec"
)

type Add struct {
}

func (a *Add) Equals(source string) bool {
	return source == "a" || source == "add"
}

func (a *Add) Exec(command []string) {
	cmd := exec.Command("git", append([]string{"add"}, command[1:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
