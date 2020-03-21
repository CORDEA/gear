package main

import (
	"os"
	"os/exec"
)

type Rebase struct {
}

func (r *Rebase) Equals(source string) bool {
	return source == "r" || source == "rebase"
}

func (r *Rebase) Exec(command []string) {
	cmd := exec.Command("git", append([]string{"rebase"}, command[1:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
