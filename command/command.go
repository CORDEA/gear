package command

type Command interface {
	Equals(source string) bool
	Exec(command []string)
}
