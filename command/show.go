package command

import (
	"strconv"
	"strings"
)

type Show struct {
}

func (s *Show) Equals(source string) bool {
	return source[0] == 'w'
}

func (s *Show) Exec(command []string) {
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
		ExecGitCommand(append([]string{"show"}, command[1:]...))
	} else {
		ExecGitCommand(append([]string{"show", "@~" + strconv.Itoa(history)}, command[1:]...))
	}
}
