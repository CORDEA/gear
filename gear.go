package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseCommand(cmd string) []string {
	return strings.Split(strings.TrimSpace(cmd), " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := [...]Command{
		&Commit{},
		&Add{},
		&Rebase{},
		&Status{},
	}
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		parsed := parseCommand(text)
		for _, cmd := range commands {
			if cmd.equals(parsed[0]) {
				cmd.exec(parsed[1:])
				continue
			}
		}
	}
}
