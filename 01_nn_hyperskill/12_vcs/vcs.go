package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	vcsDir     = "vcs"
	configFile = "config.txt"
	indexFile  = "index.txt"
)

var commands = []struct {
	name string
	desc string
}{
	{"config", "Get and set a username."},
	{"add", "Add a file to the index."},
	{"log", "Show commit logs."},
	{"commit", "Save changes."},
	{"checkout", "Restore a file."},
}

func initVCS() {
	err := os.MkdirAll(vcsDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filepath.Join(vcsDir, filename))
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(data))
}

func writeFile(filename, content string) error {
	return ioutil.WriteFile(filepath.Join(vcsDir, filename), []byte(content), 0644)
}

func handleConfig(args []string) {
	if len(args) == 0 {
		username := readFile(configFile)
		if username == "" {
			fmt.Println("Please, tell me who you are.")
		} else {
			fmt.Printf("The username is %s.\n", username)
		}
	} else {
		username := strings.Join(args, " ")
		err := writeFile(configFile, username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The username is %s.\n", username)
	}
}

func handleAdd(args []string) {
	if len(args) == 0 {
		trackedFiles := readFile(indexFile)
		if trackedFiles == "" {
			fmt.Println("Add a file to the index.")
		} else {
			fmt.Println("Tracked files:")
			fmt.Println(trackedFiles)
		}
	} else {
		filename := args[0]
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Printf("Can't find '%s'.\n", filename)
			return
		}

		trackedFiles := readFile(indexFile)
		if !strings.Contains(trackedFiles, filename) {
			if trackedFiles != "" {
				trackedFiles += "\n"
			}
			trackedFiles += filename
			err := writeFile(indexFile, trackedFiles)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("The file '%s' is tracked.\n", filename)
	}
}

func printHelp() {
	fmt.Println("These are SVCS commands:")
	for _, cmd := range commands {
		fmt.Printf("%-10s %s\n", cmd.name, cmd.desc)
	}
}

func main() {
	initVCS()

	if len(os.Args) < 2 || os.Args[1] == "--help" {
		printHelp()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "config":
		handleConfig(args)
	case "add":
		handleAdd(args)
	default:
		found := false
		for _, cmd := range commands {
			if cmd.name == command {
				fmt.Println(cmd.desc)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("'%s' is not a SVCS command.\n", command)
		}
	}
}
