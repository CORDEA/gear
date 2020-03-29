package command

import (
	"log"
	"strconv"
)

type AddList struct {
	candidates []candidate
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
	a.candidates = []candidate{}
	for _, can := range listFiles() {
		if can.status == unstaged {
			a.candidates = append(a.candidates, can)
		}
	}
	for i, candidate := range a.candidates {
		candidate.println(i)
	}
}
