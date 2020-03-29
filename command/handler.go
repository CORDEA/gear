package command

import "strings"

type Handler struct {
	commands []Command
	history  *History
}

func NewHandler(commands []Command, history *History) Handler {
	return Handler{
		commands: commands,
		history:  history,
	}
}

func parseCommand(cmd string) []string {
	return strings.Split(cmd, " ")
}

func (h *Handler) Handle(rawCmd string) {
	trimmed := strings.TrimSpace(rawCmd)
	parsed := parseCommand(trimmed)
	for _, cmd := range h.commands {
		if cmd.Equals(parsed[0]) {
			cmd.Exec(parsed)
			h.history.AddHistory(trimmed)
			continue
		}
	}
}
