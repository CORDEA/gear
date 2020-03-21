package main

import (
	"os"
	"os/exec"
)

type Diff struct {
}

func (d *Diff) Equals(source string) bool {
	return source == "d" || source == "diff"
}

func (d *Diff) Exec(command []string) {
	cmd := exec.Command("git", append([]string{"diff"}, command[1:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
