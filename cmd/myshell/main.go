package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// function that checks if a mentioned command is built-in or not
func isBuiltIn(value string) bool {
	commands := []string{
		"exit",
		"echo",
		"type",
	}

	// iterating through the list of known built-in commands
	for i := range commands {
		if commands[i] == value {
			return true
		}
	}

	return false
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		
		// Wait for user input
		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			os.Exit(1)
		}

		message = strings.TrimSpace(message)

		command := strings.Split(message, " ")

		switch command[0] {
		case "exit":
			if len(command) > 1 {
				code, err := strconv.Atoi(command[1])
	
				if err != nil {
					fmt.Fprintf(os.Stdout, "%s: command not found\n", message)
					os.Exit(1)
				}
				os.Exit(code)
			} else {
				fmt.Fprintf(os.Stdout, "Mention exit code\n")
			}
		
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(command[1:], " "))
		
		case "type":
			value := strings.Join(command[1:], " ")
			if isBuiltIn(value) {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", value)
			} else {
				fmt.Fprintf(os.Stdout, "%s: not found\n", value)
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", message)
		}

	}
}
