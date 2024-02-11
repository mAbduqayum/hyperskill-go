package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Reader
var out *bufio.Writer

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
	var s string
	fmt.Fscan(in, &s)
	res := isValid(s)
	if res {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}
	prevChar := int32(s[0])
	if prevChar != 'M' {
		return false
	}
	for _, char := range s[1:] {
		if char == 'M' {
			if prevChar != 'C' && prevChar != 'D' {
				return false
			}
		} else if char == 'R' {
			if prevChar != 'M' {
				return false
			}
		} else if char == 'C' {
			if prevChar != 'M' && prevChar != 'R' {
				return false
			}
		} else if char == 'D' {
			if prevChar != 'M' && prevChar != 'R' {
				return false
			}
		}
		prevChar = char
	}
	lastChar := prevChar
	return lastChar == 'D'
}
