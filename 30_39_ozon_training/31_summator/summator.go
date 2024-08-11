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
	for i := 1; i <= t; i++ {
		processTest()
	}
}

func processTest() {
	var a, b int
	fmt.Fscan(in, &a, &b)
	fmt.Fprintln(out, a+b)
}
