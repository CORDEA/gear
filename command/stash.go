package command

type Stash struct {
}

func (s *Stash) Equals(source string) bool {
	return source == "st" || source == "stash"
}

func (s *Stash) Exec(command []string) {
	ExecGitCommand(append([]string{"stash"}, command[1:]...))
}
