package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	ShellBuiltin CommandType = iota
	ExternalCommand
	CommandNotFound
	EmptyCommand
)

type CommandType int

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
			cmdType, cmdPath := getCommandTypeAndPath(cmdArgs[0])
			switch cmdType {
			case ShellBuiltin:
				fmt.Printf("%s is a shell builtin\n", cmdArgs[0])
			case ExternalCommand:
				fmt.Printf("%s is %s\n", cmdArgs[0], cmdPath)
			case CommandNotFound:
				fmt.Printf("%s: not found\n", cmdArgs[0])
			default:
				fmt.Println()
			}
		default:
			cmdType, cmdPath := getCommandTypeAndPath(cmd)
			switch cmdType {
			case ExternalCommand:
				cmdExec := exec.Command(cmdPath, cmdArgs...)
				cmdExec.Stdout = os.Stdout
				cmdExec.Stderr = os.Stderr
				err := cmdExec.Run()
				if err != nil {
					fmt.Println(err)
				}
				case CommandNotFound:
					fmt.Printf("%s: command not found\n", cmd)
			}
		}
	}
}

func getCommandTypeAndPath(cmd string) (CommandType, string) {
	if cmd == "" {
		return EmptyCommand, ""
	}
	isShellBuiltin := checkShellBuiltin(cmd)
	if isShellBuiltin {
		return ShellBuiltin, ""
	}
	if path, err := exec.LookPath(cmd); err == nil {
		return ExternalCommand, path
	}
	return CommandNotFound, ""
}

func checkShellBuiltin(cmd string) bool {
	return cmd == "echo" || cmd == "exit" || cmd == "type"
}
