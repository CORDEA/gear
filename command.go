package main

type Command interface {
	equals(source string) bool
	exec(args []string)
}
