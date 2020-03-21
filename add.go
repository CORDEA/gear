package main

type Add struct {
}

func (a *Add) Equals(source string) bool {
	return source == "a" || source == "add"
}

func (a *Add) Exec(command []string) {
	ExecGitCommand(append([]string{"add"}, command[1:]...))
}
