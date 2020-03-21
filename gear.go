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
			if cmd.Equals(parsed[0]) {
				cmd.Exec(parsed)
				history.AddHistory(trimmed)
				continue
			}
		}
	}
}
