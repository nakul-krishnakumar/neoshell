package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
			code, err := strconv.Atoi(command[1])

			if err != nil {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", message)
				os.Exit(1)
			}
			os.Exit(code)
		
		case "echo":
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(command[1:], " "))
		
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", message)
		}

		// exit 0 -> command to exit the shell
		// if command == "exit 0\n" {
		// 	os.Exit(0)
		// }

		// if strings.HasPrefix(command, "echo") {
		// 	fmt.Println(strings.TrimLeft(command[4:len(command)-1], " "))
		// } else {
		// 	fmt.Println(command[:len(command)-1] + ": command not found")
		// }

	}
}
