package command

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type candidate struct {
	mode   string
	status status
	file   string
}

type status int

const (
	_ status = iota
	staged
	unstaged
)

func listFiles() []candidate {
	cmd := exec.Command("git", "status", "-s")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
	result := string(buffer.Bytes())
	var candidates []candidate
	for _, line := range strings.Split(result, "\n") {
		if len(line) < 1 {
			continue
		}
		rawStatus := line[:2]
		file := line[3:]
		candidates = append(candidates, candidate{
			mode:   rawStatus,
			status: detectStatus(rawStatus),
			file:   file,
		})
	}
	return candidates
}

func detectStatus(mode string) status {
	if mode == "??" || mode == "!!" {
		return unstaged
	}
	if mode[0] != ' ' {
		return staged
	}
	return unstaged
}

func (c *candidate) println(index int) {
	fmt.Println(index, ":", c.mode, c.file)
}
