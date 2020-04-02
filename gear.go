package main

import (
	"fmt"
	"gear/command"
	"github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()
	history := command.NewHistory()
	defer history.Close()
	handler := command.NewHandler(
		[]command.Command{
			&command.Commit{},
			&command.CommitAmend{},
			&command.CommitMessage{},
			&command.CherryPick{},
			&command.Add{},
			&command.AddList{},
			&command.Rebase{},
			&command.Reset{},
			&command.RestoreList{},
			&command.Status{},
			&command.Stash{},
			&command.StashApply{},
			&history,
			&command.Diff{},
			&command.Show{},
			&command.Log{},
			&command.Checkout{},
			&command.Branch{},
			&command.CheckoutBranch{},
			&command.Quit{},
		},
		&history,
	)
	scroller := command.NewScroller(&history)
	var cmd string
	entered := true
	for {
		if entered {
			fmt.Print("> ")
			entered = false
		}
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
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
			case termbox.KeySpace:
				cmd += " "
				fmt.Print(" ")
			default:
				ch := string(ev.Ch)
				cmd += ch
				fmt.Print(ch)
			}
		}
	}

}
