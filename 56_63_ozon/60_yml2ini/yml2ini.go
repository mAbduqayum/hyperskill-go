package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solveTestCase()
		fmt.Fprintln(out)
	}
}

func solveTestCase() {
	var n int
	fmt.Fscan(in, &n)
	in.ReadString('\n') // consume newline

	yaml := make([]string, n)
	for i := 0; i < n; i++ {
		line, _ := in.ReadString('\n')
		yaml[i] = strings.TrimRight(line, " \r\n")
	}

	convertYAMLtoINI(yaml)
}

func convertYAMLtoINI(yaml []string) {
	var path []string
	var lastPath string
	var output []string
	lastIndentLevel := 0

	for _, line := range yaml {
		indentLevel := countIndentLevel(line)
		line = strings.TrimSpace(line)

		// Adjust `path` based on indentation change
		if indentLevel < lastIndentLevel {
			levelChange := lastIndentLevel - indentLevel
			path = path[:len(path)-levelChange]
		}
		lastIndentLevel = indentLevel

		if strings.HasSuffix(line, ":") {
			// This is a section
			section := line[:len(line)-1]
			path = append(path, section)
		} else {
			// This is a key-value pair
			parts := strings.Split(line, ": ")
			key, value := parts[0], parts[1]

			currentPath := ""
			if len(path) > 0 {
				currentPath = "[" + strings.Join(path, ".") + "]"
			}

			if currentPath != lastPath {
				if len(output) > 0 {
					output = append(output, "")
				}
				if currentPath != "" {
					output = append(output, currentPath)
				}
				lastPath = currentPath
			}

			output = append(output, fmt.Sprintf("%s = %s", key, value))
		}
	}

	fmt.Fprint(out, strings.Join(output, "\n"), "\n")
}

func countIndentLevel(s string) int {
	leftSpacesCount := len(s) - len(strings.TrimLeft(s, " "))
	return leftSpacesCount / 4
}
