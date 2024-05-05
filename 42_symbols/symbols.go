package main

import (
	"bufio"
	"fmt"
	"os"
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
		processTest()
	}
}

func processTest() {
	var s string
	fmt.Fscan(in, &s)
	if len(s) == 1 {
		fmt.Fprintln(out, "YES")
		return
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			if i == len(s)-2 || s[i] != s[i+2] {
				fmt.Fprintln(out, "NO")
				return
			}
			i++
		}
	}
	fmt.Fprintln(out, "YES")
}
