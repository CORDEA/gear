package main

import (
	"os"
	"os/exec"
)

type Commit struct {
}

func (c *Commit) Equals(source string) bool {
	return source == "c" || source == "commit"
}

func (c *Commit) Exec(command []string) {
	cmd := exec.Command("git", append([]string{"commit"}, command[1:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
