package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		cmd, cmdArgs := cmdFields[0], cmdFields[1:]

		switch cmd {
		case "exit":
			return
		case "echo":
			fmt.Println(strings.Join(cmdArgs, " "))
		case "type":
			getCommandType(cmdArgs)
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

func getCommandType(cmdArgs []string) {
	if (len(cmdArgs) < 1) {
		return
	}

	cmd := cmdArgs[0]
	isShellBuiltin := checkShellBuiltin(cmd)
	if isShellBuiltin {
		fmt.Printf("%s is a shell builtin\n", cmd)
	} else if path, err := exec.LookPath(cmd); err == nil {
		fmt.Printf("%s is %s\n", cmd, path)
	} else {
		fmt.Printf("%s: command not found\n", cmd)
	}
}

func checkShellBuiltin(cmd string) bool {
	return cmd == "echo" || cmd == "exit" || cmd == "type"
}