package main

import (
	"log"
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
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalln("Failed to execute command. ", err)
	}
}
