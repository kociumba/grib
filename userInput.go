package main

import (
	"bufio"
	"os"
)

func hold(exit chan bool) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	if text == "exit\n" {
		exit <- true
	}
}
