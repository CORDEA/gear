package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
)

type command interface {
	equals(source string) bool
	exec(args []string)
}

type commit struct {
}

func (c *commit) equals(source string) bool {
	return source == "c"
}

func (c *commit) exec(args []string) {
	cmd := exec.Command("git", append([]string{"commit"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalln("Failed to execute command. ", err)
	}
}

func parseCommand(cmd string) []string {
	return strings.Split(strings.TrimSpace(cmd), " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := [...]command{
		&commit{},
	}
	for {
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
