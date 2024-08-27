package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	vcs := NewVCS()

	if len(os.Args) < 2 || os.Args[1] == "--help" {
		vcs.PrintHelp()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "config":
		vcs.HandleConfig(args)
	case "add":
		vcs.HandleAdd(args)
	case "commit":
		vcs.HandleCommit(args)
	case "log":
		vcs.HandleLog()
	case "checkout":
		vcs.HandleCheckout(args)
	default:
		if !vcs.HandleUnknownCommand(command) {
			fmt.Printf("'%s' is not a SVCS command.\n", command)
		}
	}
}

const (
	vcsDir     = "vcs"
	configFile = "config.txt"
	indexFile  = "index.txt"
	logFile    = "log.txt"
	commitsDir = "commits"
)

type Command struct {
	Name string
	Desc string
}

type VCS struct {
	Commands []Command
}

func NewVCS() *VCS {
	vcs := &VCS{
		Commands: []Command{
			{"config", "Get and set a username."},
			{"add", "Add a file to the index."},
			{"log", "Show commit logs."},
			{"commit", "Save changes."},
			{"checkout", "Restore a file."},
		},
	}
	vcs.initVCS()
	return vcs
}

func (v *VCS) initVCS() {
	if err := os.MkdirAll(vcsDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func (v *VCS) readFile(filename string) string {
	data, err := os.ReadFile(filepath.Join(vcsDir, filename))
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(data))
}

func (v *VCS) writeFile(filename, content string) error {
	return os.WriteFile(filepath.Join(vcsDir, filename), []byte(content), 0644)
}

func (v *VCS) HandleConfig(args []string) {
	if len(args) == 0 {
		username := v.readFile(configFile)
		if username == "" {
			fmt.Println("Please, tell me who you are.")
		} else {
			fmt.Printf("The username is %s.\n", username)
		}
	} else {
		username := strings.Join(args, " ")
		if err := v.writeFile(configFile, username); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The username is %s.\n", username)
	}
}

func (v *VCS) HandleAdd(args []string) {
	if len(args) == 0 {
		trackedFiles := v.readFile(indexFile)
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

		trackedFiles := v.readFile(indexFile)
		if !strings.Contains(trackedFiles, filename) {
			if trackedFiles != "" {
				trackedFiles += "\n"
			}
			trackedFiles += filename
			if err := v.writeFile(indexFile, trackedFiles); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("The file '%s' is tracked.\n", filename)
	}
}

func (v *VCS) HandleCommit(args []string) {
	if len(args) == 0 {
		fmt.Println("Message was not passed.")
		return
	}

	message := strings.Join(args, " ")
	trackedFiles := strings.Split(v.readFile(indexFile), "\n")
	if len(trackedFiles) == 0 || (len(trackedFiles) == 1 && trackedFiles[0] == "") {
		fmt.Println("Nothing to commit.")
		return
	}

	lastCommitHash := v.getLastCommitHash()
	currentHash := v.calculateFilesHash(trackedFiles)

	if currentHash == lastCommitHash {
		fmt.Println("Nothing to commit.")
		return
	}

	commitDir := filepath.Join(vcsDir, commitsDir, currentHash)
	if err := os.MkdirAll(commitDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	for _, file := range trackedFiles {
		if file == "" {
			continue
		}
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(commitDir, filepath.Base(file)), content, 0644); err != nil {
			log.Fatal(err)
		}
	}

	username := v.readFile(configFile)
	logEntry := fmt.Sprintf("commit %s\nAuthor: %s\n%s", currentHash, username, message)
	v.prependToFile(filepath.Join(vcsDir, logFile), logEntry)

	fmt.Println("Changes are committed.")
}

func (v *VCS) HandleLog() {
	logContent := v.readFile(logFile)
	if logContent == "" {
		fmt.Println("No commits yet.")
	} else {
		commits := strings.Split(strings.TrimSpace(logContent), "\n\n")
		for i := len(commits) - 1; i >= 0; i-- {
			fmt.Println(commits[i])
			if i > 0 {
				fmt.Println()
			}
		}
	}
}

func (v *VCS) HandleCheckout(args []string) {
	if len(args) == 0 {
		fmt.Println("Commit id was not passed.")
		return
	}

	commitID := args[0]
	commitDir := filepath.Join(vcsDir, commitsDir, commitID)

	if _, err := os.Stat(commitDir); os.IsNotExist(err) {
		fmt.Println("Commit does not exist.")
		return
	}

	trackedFiles := strings.Split(v.readFile(indexFile), "\n")
	for _, file := range trackedFiles {
		if file == "" {
			continue
		}

		sourceFile := filepath.Join(commitDir, filepath.Base(file))
		if err := v.copyFile(sourceFile, file); err != nil {
			log.Printf("Error restoring file %s: %v", file, err)
		}
	}

	fmt.Printf("Switched to commit %s.\n", commitID)
}

func (v *VCS) copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func (v *VCS) getLastCommitHash() string {
	entries, err := os.ReadDir(filepath.Join(vcsDir, commitsDir))
	if err != nil {
		return ""
	}
	if len(entries) == 0 {
		return ""
	}
	return entries[len(entries)-1].Name()
}

func (v *VCS) calculateFilesHash(files []string) string {
	h := sha256.New()
	for _, file := range files {
		if file == "" {
			continue
		}
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		h.Write(content)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (v *VCS) prependToFile(filename, content string) {
	existingContent, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	newContent := []byte(content)
	if len(existingContent) > 0 {
		newContent = append(newContent, '\n')
		newContent = append(newContent, existingContent...)
	}

	if err := os.WriteFile(filename, newContent, 0644); err != nil {
		log.Fatal(err)
	}
}

func (v *VCS) PrintHelp() {
	fmt.Println("These are SVCS commands:")
	for _, cmd := range v.Commands {
		fmt.Printf("%-10s %s\n", cmd.Name, cmd.Desc)
	}
}

func (v *VCS) HandleUnknownCommand(command string) bool {
	for _, cmd := range v.Commands {
		if cmd.Name == command {
			fmt.Println(cmd.Desc)
			return true
		}
	}
	return false
}
