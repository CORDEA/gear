package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseCommand(cmd string) []string {
	return strings.Split(cmd, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	history := NewHistory()
	defer history.Close()
	commands := [...]Command{
		&Commit{},
		&Add{},
		&Rebase{},
		&Status{},
		&history,
	}
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		trimmed := strings.TrimSpace(text)
		parsed := parseCommand(trimmed)
		for _, cmd := range commands {
			if cmd.equals(parsed[0]) {
				cmd.exec(parsed[1:])
				history.AddHistory(trimmed)
				continue
			}
		}
	}
}
