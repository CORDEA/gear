package command

type CherryPick struct {
}

func (c *CherryPick) Equals(source string) bool {
	return source == "pick" || source == "cherry-pick"
}

func (c *CherryPick) Exec(command []string) {
	ExecGitCommand(append([]string{"cherry-pick"}, command[1:]...))
}
