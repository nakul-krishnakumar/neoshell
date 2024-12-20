package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
	for _, command := range commands {
		if command == value {
			return true
		}
	}

	return false
}

func commandExistsInPath(env string, arg string) {
	paths := strings.Split(env, ":")

	for _, path := range paths {
		fp := filepath.Join(path, arg)
		if _, err := os.Stat(fp); err == nil {
			fmt.Fprintf(os.Stdout, "%s is %s\n", arg, fp)
			return 
		} 
	}
	fmt.Fprintf(os.Stdout, "%s: not found\n", arg)
	return
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
			// exit <exit_code> -> exits with <exit_code>
			// if <exit_code> is not mentioned, shell asks to mention it

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
			// echo <message> -> prints the <message>
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(command[1:], " "))
		
		case "type":
			// type <command> -> checks if <command> is built-in or invalid

			value := strings.Join(command[1:], " ")
			if isBuiltIn(value) {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", value)
			} else {
				commandExistsInPath(os.Getenv("PATH"), value)
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", message)
		}

	}
}
