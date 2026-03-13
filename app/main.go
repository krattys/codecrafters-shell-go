package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Print("$ ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	cmd := strings.Fields(text)[0]
	fmt.Printf("%s: command not found\n", cmd)
}
