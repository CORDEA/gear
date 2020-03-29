package command

type Reset struct {
}

func (r *Reset) Equals(source string) bool {
	return source == "r" || source == "res" || source == "reset"
}

func (r *Reset) Exec(command []string) {
	ExecGitCommand(append([]string{"reset"}, command[1:]...))
}
