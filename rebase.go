package main

import (
	"os"
	"os/exec"
)

type Rebase struct {
}

func (r *Rebase) equals(source string) bool {
	return source == "r" || source == "rebase"
}

func (r *Rebase) exec(args []string) {
	cmd := exec.Command("git", append([]string{"rebase"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
