package main

import (
	"fmt"
	"os"
)

var commands = map[string]string{
	"config":   "Get and set a username.",
	"add":      "Add a file to the index.",
	"log":      "Show commit logs.",
	"commit":   "Save changes.",
	"checkout": "Restore a file.",
}

func printHelp() {
	fmt.Println("These are SVCS commands:")
	for cmd, desc := range commands {
		fmt.Printf("%-10s %s\n", cmd, desc)
	}
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" {
		printHelp()
		return
	}

	command := os.Args[1]
	if desc, exists := commands[command]; exists {
		fmt.Println(desc)
	} else {
		fmt.Printf("'%s' is not a SVCS command.\n", command)
	}
}
