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
	}
}

func solveTestCase() {
	var s string
	fmt.Fscan(in, &s)
	result := solve(s)
	if result == "" {
		result = "0"
	}
	fmt.Fprintln(out, result)
}

func solve(s string) string {
	var result strings.Builder
	for i := 0; i < len(s)-1; i++ {
		if s[i] >= s[i+1] {
			result.WriteByte(s[i])
		} else {
			result.WriteString(s[i+1:])
			break
		}
	}
	return result.String()
}
