package main

import (
	"os"
	"os/exec"
)

type Diff struct {
}

func (d *Diff) equals(source string) bool {
	return source == "d" || source == "diff"
}

func (d *Diff) exec(args []string) {
	cmd := exec.Command("git", append([]string{"diff"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
