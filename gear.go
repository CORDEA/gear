package main

import (
	"bufio"
	"os"
)

type command interface {
	equals(source string) bool
	exec()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := [...]command{
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
