package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in         = bufio.NewReader(os.Stdin)
	out        = bufio.NewWriter(os.Stdout)
	toRight    = "1\n1 1 R\n"
	toDown     = "1\n1 1 D\n"
	horizontal = "2\n1 1 R\n%d %d L\n"
	vertical   = "2\n1 1 D\n%d %d U\n"
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
	var n, m int
	fmt.Fscan(in, &n, &m)

	if n == 1 {
		fmt.Fprintf(out, toRight)
	} else if m == 1 {
		fmt.Fprintf(out, toDown)
	} else if n >= m {
		fmt.Fprintf(out, vertical, n, m)
	} else {
		fmt.Fprintf(out, horizontal, n, m)
	}
}
