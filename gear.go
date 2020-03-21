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
	scroller := NewScroller(&history)
	var cmd string
	entered := true
loop:
	for {
		if entered {
			fmt.Print("> ")
			entered = false
		}
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyArrowUp:
				var err error
				if cmd, err = scroller.ScrollToAbove(); err == nil {
					termbox.Sync()
					fmt.Print("> " + cmd)
				}
			case termbox.KeyArrowDown:
				var err error
				if cmd, err = scroller.ScrollToBelow(); err == nil {
					termbox.Sync()
					fmt.Print("> " + cmd)
				}
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				if cmd != "" {
					cmd = cmd[:len(cmd)-1]
					termbox.Sync()
					fmt.Print("> " + cmd)
				}
			case termbox.KeyEnter:
				fmt.Println()
				handler.Handle(cmd)
				entered = true
				cmd = ""
				scroller.Reset()
			default:
				ch := string(ev.Ch)
				cmd += ch
				fmt.Print(ch)
			}
		}
	}

}
