package command

import "os"

type Quit struct {
}

func (q *Quit) Equals(source string) bool {
	return source == "quit"
}

func (q *Quit) Exec(command []string) {
	os.Exit(0)
}
