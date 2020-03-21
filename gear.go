package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	history := NewHistory()
	defer history.Close()
	handler := NewHandler(
		[]Command{
			&Commit{},
			&Add{},
			&Rebase{},
			&Status{},
			&history,
			&Diff{},
			&Show{},
			&Log{},
			&Checkout{},
			&Branch{},
			&CheckoutBranch{},
		},
		&history,
	)
	var cmd string
loop:
	for {
		if cmd == "" {
			fmt.Print("> ")
		}
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyEnter:
				fmt.Println()
				handler.Handle(cmd)
				cmd = ""
			default:
				ch := string(ev.Ch)
				cmd += ch
				fmt.Print(ch)
			}
		}
	}

}
