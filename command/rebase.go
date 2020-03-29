package command

type Rebase struct {
}

func (r *Rebase) Equals(source string) bool {
	return source == "rb" || source == "rebase"
}

func (r *Rebase) Exec(command []string) {
	ExecGitCommand(append([]string{"rebase"}, command[1:]...))
}
