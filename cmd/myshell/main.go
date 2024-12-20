package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		"pwd",
		"cd",
	}

	// iterating through the list of known built-in commands
	for _, command := range commands {
		if command == value {
			return true
		}
	}

	return false
}

// checks if the mentioned <arg> is found anywhere in any dir
func commandExistsInPath(env string, arg string) (bool, string) {
	paths := strings.Split(env, ":")

	for _, path := range paths {
		fp := filepath.Join(path, arg)
		if _, err := os.Stat(fp); err == nil {
			return true, fp 
		} 
	}
	return false, ""
}

// exits the code with the mentioned <exitCode>
func doExit(exitCode string, message string) {
	code, err := strconv.Atoi(exitCode)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", message)
		os.Exit(1)
	}
	os.Exit(code)
}

func main() {
// 	fmt.Printf(`
//                                     $$\               $$\$$\ 
//                                     $$ |              $$ $$ |
// $$$$$$$\  $$$$$$\  $$$$$$\  $$$$$$$\$$$$$$$\  $$$$$$\ $$ $$ |
// $$  __$$\$$  __$$\$$  __$$\$$  _____$$  __$$\$$  __$$\$$ $$ |
// $$ |  $$ $$$$$$$$ $$ /  $$ \$$$$$$\ $$ |  $$ $$$$$$$$ $$ $$ |
// $$ |  $$ $$   ____$$ |  $$ |\____$$\$$ |  $$ $$   ____$$ $$ |
// $$ |  $$ \$$$$$$$\\$$$$$$  $$$$$$$  $$ |  $$ \$$$$$$$\$$ $$ |
// \__|  \__|\_______|\______/\_______/\__|  \__|\_______\__\__|
	
// `)
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
				doExit(command[1], message)
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
				exists, fp := commandExistsInPath(os.Getenv("PATH"), value)
				if exists {
					fmt.Fprintf(os.Stdout, "%s is %s\n", value, fp)
				} else {
					fmt.Fprintf(os.Stdout, "%s: not found\n", value)
				}
			}

		case "pwd":
			fp, err := os.Getwd()
			if err == nil {
				fmt.Fprintf(os.Stdout ,"%s\n", fp)
			}

		case "cd":
			if len(command) > 1 {
				path := command[1]
				if command[1] == "~" {
					path, err = os.UserHomeDir()
					if err != nil {
						fmt.Printf(os.Stdout, "Error: %s", err)
					}
				} 

				err := os.Chdir(path)
				if err != nil {
					fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", command[1])
				
				}
			} else {
				fmt.Fprintf(os.Stdout, "cd: : No such file or directory\n")
			}
		default:
			cmdExec := exec.Command(command[0], command[1:]...)
			cmdExec.Stderr = os.Stderr
			cmdExec.Stdout = os.Stdout

			err := cmdExec.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", command[0])
			}
		}
	}
}

/* 

                                    $$\               $$\$$\ 
                                    $$ |              $$ $$ |
$$$$$$$\  $$$$$$\  $$$$$$\  $$$$$$$\$$$$$$$\  $$$$$$\ $$ $$ |
$$  __$$\$$  __$$\$$  __$$\$$  _____$$  __$$\$$  __$$\$$ $$ |
$$ |  $$ $$$$$$$$ $$ /  $$ \$$$$$$\ $$ |  $$ $$$$$$$$ $$ $$ |
$$ |  $$ $$   ____$$ |  $$ |\____$$\$$ |  $$ $$   ____$$ $$ |
$$ |  $$ \$$$$$$$\\$$$$$$  $$$$$$$  $$ |  $$ \$$$$$$$\$$ $$ |
\__|  \__|\_______|\______/\_______/\__|  \__|\_______\__\__|


*/