package command

type StashApply struct {
}

func (s *StashApply) Equals(source string) bool {
	return source == "sta" || source == "stashapply"
}

func (s *StashApply) Exec(command []string) {
	ExecGitCommand(append([]string{"stash", "apply"}, command[1:]...))
}
