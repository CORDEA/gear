package main

import (
	"bufio"
	"os"
)

type command interface {
	equals(source string) bool
	exec()
}

type commit struct {
}

func (c *commit) equals(source string) bool {
	return source == "c"
}

func (c *commit) exec() {
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := [...]command{
		&commit{},
	}
	for {
		text, _:= reader.ReadString('\n')
		for _, cmd := range commands {
			if cmd.equals(text) {
				cmd.exec()
				continue
			}
		}
	}
}
