package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _:= reader.ReadString('\n')
		fmt.Println(strings.TrimSpace(text))
	}
}
