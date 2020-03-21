package main

type Commit struct {
}

func (c *Commit) Equals(source string) bool {
	return source == "c" || source == "commit"
}

func (c *Commit) Exec(command []string) {
	ExecGitCommand(append([]string{"commit"}, command[1:]...))
}
