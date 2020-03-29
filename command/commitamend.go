package command

type CommitAmend struct {
}

func (c *CommitAmend) Equals(source string) bool {
	return source == "ca" || source == "cam" || source == "commit-amend"
}

func (c *CommitAmend) Exec(command []string) {
	if command[0] == "cam" {
		ExecGitCommand(append([]string{"commit", "--amend", "-m"}, command[1:]...))
		return
	}
	ExecGitCommand(append([]string{"commit", "--amend"}, command[1:]...))
}
