package main

type Command interface {
	Equals(source string) bool
	Exec(command []string)
}
