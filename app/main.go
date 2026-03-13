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
		fmt.Print("$ ")
		text, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		text = strings.Trim(text, "\n")
		cmdFields := strings.Fields(text)
		cmd, _ := cmdFields[0], strings.Join(cmdFields[1:], " ")

		if cmd == "exit" {
			return
		} else {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
