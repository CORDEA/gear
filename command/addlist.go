package command

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type AddList struct {
	candidates []candidate
}

type candidate struct {
	mode string
	file string
}

func (a *AddList) Equals(source string) bool {
	return source == "al" || source == "add-list"
}

func (a *AddList) Exec(command []string) {
	if len(command) == 1 {
		a.listFiles()
		return
	}
	var files []string
	for _, cmd := range command[1:] {
		i, err := strconv.Atoi(cmd)
		if err != nil {
			log.Println("add-list accepts only number. ", cmd)
			continue
		}
		if i >= len(a.candidates) {
			continue
		}
		files = append(files, a.candidates[i].file)
	}
	ExecGitCommand(append([]string{"add"}, files...))
}

func (a *AddList) listFiles() {
	cmd := exec.Command("git", "status", "-s")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
	result := string(buffer.Bytes())
	a.candidates = []candidate{}
	for _, line := range strings.Split(result, "\n") {
		can := strings.Split(strings.TrimSpace(line), " ")
		if len(can) != 2 {
			continue
		}
		a.candidates = append(a.candidates, candidate{
			mode: can[0],
			file: can[1],
		})
	}
	for i, candidate := range a.candidates {
		candidate.println(i)
	}
}

func (c *candidate) println(index int) {
	fmt.Println(index, ":", c.mode, c.file)
}
