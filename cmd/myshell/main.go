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
	fmt.Printf(`
										$$\               $$\$$\ 
										$$ |              $$ $$ |
	$$$$$$$\  $$$$$$\  $$$$$$\  $$$$$$$\$$$$$$$\  $$$$$$\ $$ $$ |
	$$  __$$\$$  __$$\$$  __$$\$$  _____$$  __$$\$$  __$$\$$ $$ |
	$$ |  $$ $$$$$$$$ $$ /  $$ \$$$$$$\ $$ |  $$ $$$$$$$$ $$ $$ |
	$$ |  $$ $$   ____$$ |  $$ |\____$$\$$ |  $$ $$   ____$$ $$ |
	$$ |  $$ \$$$$$$$\\$$$$$$  $$$$$$$  $$ |  $$ \$$$$$$$\$$ $$ |
	\__|  \__|\_______|\______/\_______/\__|  \__|\_______\__\__|
		
	`)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		
		// Wait for user input
		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input: ", err)
			os.Exit(1)
		}

		// message -> user input ( command + arguments )
		message = strings.TrimSpace(message)
		messageArray := strings.Split(message, " ")

		command := messageArray[0]
		args := messageArray[1:]


		// TODO FIX QUOTING SUPPORT
		// I only though about quotes being on starting and ending, 
		// but didnt think about the possibility of multiple quoted strings
		// like -> echo "hello   " "world"
		if len(args) > 0 {
			for i, arg := range args {
				if (strings.HasPrefix(arg, `"`) && strings.HasSuffix(arg, `""`)) ||
					(strings.HasPrefix(arg, `'`) && strings.HasSuffix(arg, `'`)) {
					args[i] = strings.Trim(arg, `"'`)
				}
			}
		}
			// if strings.HasPrefix(args[0], `'`) && strings.HasSuffix(args[len(args)-1], `'`) ||
			// 	strings.HasPrefix(args[0], `"`) && strings.HasSuffix(args[len(args)-1], `"`) {
			// 		args[0] = args[0][1:]
			// 		var found bool
			// 		args[len(args)-1], found = strings.CutSuffix(args[len(args)-1], `'`)
			// 		if !found {
			// 			args[len(args)-1], _ = strings.CutSuffix(args[len(args)-1], `"`)
			// 		}
			// 	} 
			// }

		switch command {
		case "exit":
			// exit <exit_code> -> exits with <exit_code>
			// if <exit_code> is not mentioned, shell asks to mention it

			if len(command) > 1 {
				// args[0] -> exitCode string
				// message -> message string ( full user input to print when err occurs )
				doExit(args[0], message) 
			} else {
				fmt.Fprintf(os.Stdout, "Mention exit code\n")
			}
		
		case "echo":
			// echo <message> -> prints the <message>
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
		
		case "type":
			// type <command> -> checks if <command> is built-in or invalid

			value := strings.Join(args, " ")
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
			// pwd -> prints the PRESENT WORKING DIRECTORY
			fp, err := os.Getwd()
			if err == nil {
				fmt.Fprintf(os.Stdout ,"%s\n", fp)
			}

		case "cd":
			// cd <location> -> navigates to the given <location>

			if len(args) >= 1 {
				path := args[0]
				if path == "~" {
					path, err = os.UserHomeDir()
					if err != nil {
						fmt.Fprintf(os.Stdout, "Error: %s", err)
					}
				} 

				err := os.Chdir(path)
				if err != nil {
					fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", args[0])
				
				}
			} else {
				fmt.Fprintf(os.Stdout, "cd: : No such file or directory\n")
			}
		default:
			// executes default command line's commands
			cmdExec := exec.Command(command, args...)
			cmdExec.Stderr = os.Stderr
			cmdExec.Stdout = os.Stdout

			err := cmdExec.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", command)
			}
		}
	}
}