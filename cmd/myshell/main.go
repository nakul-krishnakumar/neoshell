package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i:=0; ;i++ {
		fmt.Fprint(os.Stdout, "$ ")
		
		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			os.Exit(1)
		}

		// exit 0 -> command to exit the shell
		if command == "exit 0\n" {
			os.Exit(0)
		}

		if strings.HasPrefix(command, "echo") {
			fmt.Println(strings.TrimLeft(command[4:len(command)-1], " "))
		} else {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}

	}
}
