package main

import (
	"log"
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
	if err := cmd.Run(); err != nil {
		log.Fatalln("Failed to execute command. ", err)
	}
}
