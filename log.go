package main

type Log struct {
}

func (l *Log) Equals(source string) bool {
	return source == "l" || source == "log"
}

func (l *Log) Exec(command []string) {
	ExecGitCommand(append([]string{"log"}, command[1:]...))
}
