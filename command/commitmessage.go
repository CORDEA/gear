package command

type CommitMessage struct {
}

func (c *CommitMessage) Equals(source string) bool {
	return source == "cm" || source == "commit-message"
}

func (c *CommitMessage) Exec(command []string) {
	ExecGitCommand(append([]string{"commit", "-m"}, command[1:]...))
}
