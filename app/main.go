package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Print the prompt
		fmt.Fprint(os.Stdout, "$ ")

		// Read user input
		command, err := reader.ReadString('\n')
		if err != nil {
			// If EOF is reached or input error, exit gracefully
			break
		}

		// Remove trailing newline character
		command = strings.TrimSpace(command)
		if command == "exit 0" {
			os.Exit(0)
			break
		}
		if strings.HasPrefix(command, "echo") {
			parts := strings.Fields(command)
			if len(parts) > 1 {
				output := strings.Join(parts[1:], " ")
				fmt.Println(output)
			} else {
				fmt.Println(" ")
			}
			continue
		}
		if strings.HasPrefix(command, "type") {
			parts := strings.Fields(command)
			if len(parts) > 1 {
				if parts[len(parts)-1] == "echo" || parts[len(parts)-1] == "exit" || parts[len(parts)-1] == "type" {
					fmt.Printf("%s: is a shell builtin\n", parts[len(parts)-1])

				} else {
					fmt.Printf("%s: not found\n", parts[len(parts)-1])
				}
			} else {
				fmt.Println(" ")
			}
			continue
		}
		// Skip empty input
		if command == "" {
			continue
		}

		// Print command not found
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}
