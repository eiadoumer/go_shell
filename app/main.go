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

		// Skip empty input
		if command == "" {
			continue
		}

		// Print command not found
		fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
	}
}
