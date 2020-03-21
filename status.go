package main

import (
	"os"
	"os/exec"
)

type Status struct {
}

func (s *Status) Equals(source string) bool {
	return source == "s" || source == "status"
}

func (s *Status) Exec(command []string) {
	cmd := exec.Command("git", append([]string{"status"}, command[1:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
