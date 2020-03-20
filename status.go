package main

import (
	"os"
	"os/exec"
)

type Status struct {
}

func (s *Status) equals(source string) bool {
	return source == "s" || source == "status"
}

func (s *Status) exec(args []string) {
	cmd := exec.Command("git", append([]string{"status"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
