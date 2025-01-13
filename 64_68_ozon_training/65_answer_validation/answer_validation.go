package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	var n int
	fmt.Fscan(in, &n)
	expected := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Fscan(in, &expected[i])
		if err != nil {
			fmt.Fprintln(out, "no")
			return
		}
	}

	_, _ = in.ReadString('\n')
	line, _ := in.ReadString('\n')
	if line[0] == ' ' || line[len(line)-1] == ' ' || strings.Contains(line, "  ") {
		fmt.Fprintln(out, "no")
		return
	}
	output := strings.Fields(line)
	if len(output) != n {
		fmt.Fprintln(out, "no")
		return
	}
	actual := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Sscanf(output[i], "%d", &actual[i])
		if err != nil {
			fmt.Fprintln(out, "no")
			return
		}
	}
	slices.Sort(expected)
	if slices.Equal(expected, actual) {
		fmt.Fprintln(out, "yes")
	} else {
		fmt.Fprintln(out, "no")
	}
}
