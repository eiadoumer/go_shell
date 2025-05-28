package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
				fmt.Println()
			}
			continue
		}
		if strings.HasPrefix(command, "type") {
			parts := strings.Fields(command)
			if len(parts) > 1 {
				cmdName := parts[1] // Get the command name to search for

				// Check if it's a builtin first
				if cmdName == "echo" || cmdName == "exit" || cmdName == "type" {
					fmt.Printf("%s is a shell builtin\n", cmdName)
				} else {
					// Search in PATH
					pathFound := findInPath(cmdName)
					if pathFound != "" {
						fmt.Printf("%s is %s\n", cmdName, pathFound)
					} else {
						fmt.Printf("%s: not found\n", cmdName)
					}
				}
			}
			continue
		}

		// Skip empty input
		if command == "" {
			continue
		}

		// Try to execute as external program
		parts := strings.Fields(command)
		if len(parts) > 0 {
			cmdName := parts[0]
			args := parts[1:] // Get arguments (everything after the command name)

			// Find the command in PATH
			cmdPath := findInPath(cmdName)
			if cmdPath != "" {
				// Execute the external program
				// Create arguments array with command name as first argument
				cmd := exec.Command(cmdPath, args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				err := cmd.Run()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error executing %s: %v\n", cmdName, err)
				}
			} else {
				// Print command not found
				fmt.Fprintf(os.Stdout, "%s: command not found\n", cmdName)
			}
		}
	}
}

// findInPath searches for a command in the directories listed in PATH
func findInPath(cmdName string) string {
	// Get the PATH environment variable
	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return ""
	}

	// Split PATH by colons to get individual directories
	pathDirs := strings.Split(pathEnv, ":")

	// Search each directory in order
	for _, dir := range pathDirs {
		if dir == "" {
			continue
		}

		// Create the full path to the potential executable
		fullPath := filepath.Join(dir, cmdName)

		// Check if the file exists and is executable
		if fileInfo, err := os.Stat(fullPath); err == nil {
			// Check if it's a regular file and executable
			if fileInfo.Mode().IsRegular() && (fileInfo.Mode().Perm()&0111) != 0 {
				return fullPath
			}
		}
	}

	return "" // Not found
}
