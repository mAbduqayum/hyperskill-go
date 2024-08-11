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
	processTest()
}

func processTest() {
	var x string
	fmt.Fscanln(in, &x)
	fmt.Fprintln(out, x)
}
