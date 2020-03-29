package command

type CheckoutBranch struct {
}

func (c *CheckoutBranch) Equals(source string) bool {
	return source == "cb" || source == "bb"
}

func (c *CheckoutBranch) Exec(command []string) {
	ExecGitCommand(append([]string{"checkout", "-b"}, command[1:]...))
}
