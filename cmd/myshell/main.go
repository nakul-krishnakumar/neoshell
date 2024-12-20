package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	userCommand, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	strippedUserCommand := strings.ReplaceAll(userCommand, "\n", "")
	
	if strippedUserCommand == "invalid_command" {
		fmt.Fprint(os.Stdout, strippedUserCommand + ": command not found")
	}
	
}
