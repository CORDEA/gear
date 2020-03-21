package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type History struct {
	histories  []string
	historyMap map[int]string
	file       *os.File
}

func NewHistory() History {
	file, err := os.OpenFile(filepath.Join(os.Getenv("HOME"), ".gear_history"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	var histories []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		histories = append(histories, scanner.Text())
	}
	return History{histories: histories, file: file}
}

func (h *History) Equals(source string) bool {
	return source == "h" || source == "history"
}

func (h *History) Exec(command []string) {
	query := strings.Join(command[1:], " ")
	h.historyMap = map[int]string{}
	var n = 0
	for i := len(h.histories) - 1; i >= 0; i-- {
		history := h.histories[i]
		if strings.Contains(history, query) {
			fmt.Println(n, ": ", history)
			h.historyMap[n] = history
			n++
		}
	}
}

func (h *History) AddHistory(cmd string) {
	for _, history := range h.histories {
		if history == cmd {
			return
		}
	}
	h.histories = append(h.histories, cmd)
	if _, err := h.file.WriteString(cmd + "\n"); err != nil {
		log.Fatalln(err)
	}
}

func (h *History) Close() {
	if err := h.file.Close(); err != nil {
		log.Fatalln(err)
	}
}
