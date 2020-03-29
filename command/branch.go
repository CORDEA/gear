package command

type Branch struct {
}

func (b *Branch) Equals(source string) bool {
	return source == "b" || source == "branch"
}

func (b *Branch) Exec(command []string) {
	ExecGitCommand(append([]string{"branch"}, command[1:]...))
}
