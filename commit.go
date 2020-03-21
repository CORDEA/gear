package main

import (
	"os"
	"os/exec"
)

type Commit struct {
}

func (c *Commit) equals(source string) bool {
	return source == "c" || source == "commit"
}

func (c *Commit) exec(args []string) {
	cmd := exec.Command("git", append([]string{"commit"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
