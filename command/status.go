package command

type Status struct {
}

func (s *Status) Equals(source string) bool {
	return source == "s" || source == "status"
}

func (s *Status) Exec(command []string) {
	ExecGitCommand(append([]string{"status"}, command[1:]...))
}
