package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var in *bufio.Reader
var out *bufio.Writer

type Folder struct {
	Name    string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func (f Folder) FilesCount() int {
	count := len(f.Files)
	for _, subFolder := range f.Folders {
		count += subFolder.FilesCount()
	}
	return count
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		processTest()
	}
}

func processTest() {
	var n int
	fmt.Fscan(in, &n)
	in.ReadString('\n')
	rows := make([]string, n)
	for i := 0; i < n; i++ {
		rows[i], _ = in.ReadString('\n')
	}
	rawSting := strings.Join(rows, "")
	var folder Folder
	err := json.Unmarshal([]byte(rawSting), &folder)
	if err != nil {
		fmt.Fprintf(out, "Error parsing JSON: %v\n", err)
		return
	}
	fmt.Fprintln(out, countVirus(&folder))
}

func countVirus(f *Folder) int {
	count := 0
	hasVirus := false
	for _, file := range f.Files {
		if strings.HasSuffix(file, ".hack") {
			hasVirus = true
			break
		}
	}
	if hasVirus {
		count += f.FilesCount()
		return count
	}
	for _, subFolder := range f.Folders {
		count += countVirus(&subFolder)
	}
	return count
}
