package command

import (
	"strconv"
	"strings"
)

type Diff struct {
}

func (d *Diff) Equals(source string) bool {
	if len(source) < 1 {
		return false
	}
	return source[0] == 'd'
}

func (d *Diff) Exec(command []string) {
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
		ExecGitCommand(append([]string{"diff"}, command[1:]...))
	} else {
		ExecGitCommand(append([]string{"diff", "@~" + strconv.Itoa(history)}, command[1:]...))
	}
}
