package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Diff struct {
}

func (d *Diff) Equals(source string) bool {
	return source[0] == 'd'
}

func (d *Diff) Exec(command []string) {
	var cmd *exec.Cmd
	existsHead := false
	for _, cmd := range command[1:] {
		if strings.Contains(cmd, "HEAD") {
			existsHead = true
		}
		if strings.Contains(cmd, "@") {
			existsHead = true
		}
	}
	history := len(command[0]) - 1
	if existsHead || history == 0 {
		cmd = exec.Command("git", append([]string{"diff"}, command[1:]...)...)
	} else {
		cmd = exec.Command("git", append([]string{"diff", "@~" + strconv.Itoa(history)}, command[1:]...)...)
	}
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
