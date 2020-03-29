package command

import (
	"log"
	"strconv"
)

type RestoreList struct {
	candidates []candidate
}

func (r *RestoreList) Equals(source string) bool {
	return source == "r@" || source == "r!" || source == "restore-list"
}

func (r *RestoreList) Exec(command []string) {
	if len(command) == 1 {
		r.listFiles()
		return
	}
	var files []string
	for _, cmd := range command[1:] {
		i, err := strconv.Atoi(cmd)
		if err != nil {
			log.Println("restore-list accepts only number. ", cmd)
			continue
		}
		if i >= len(r.candidates) {
			continue
		}
		files = append(files, r.candidates[i].file)
	}
	ExecGitCommand(append([]string{"restore", "--staged"}, command[1:]...))
}

func (r *RestoreList) listFiles() {
	r.candidates = []candidate{}
	for _, can := range listFiles() {
		if can.status == staged {
			r.candidates = append(r.candidates, can)
		}
	}
	for i, candidate := range r.candidates {
		candidate.println(i)
	}
}
