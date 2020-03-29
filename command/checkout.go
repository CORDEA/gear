package command

type Checkout struct {
}

func (c *Checkout) Equals(source string) bool {
	return source == "ch" || source == "checkout"
}

func (c *Checkout) Exec(command []string) {
	ExecGitCommand(append([]string{"checkout"}, command[1:]...))
}
