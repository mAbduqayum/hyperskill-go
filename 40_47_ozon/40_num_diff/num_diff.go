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
	var a, b int
	fmt.Fscan(in, &a, &b)
	fmt.Fprintln(out, a-b)
}
