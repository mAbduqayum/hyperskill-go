package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

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

	l := make([]int, n)
	r := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &r[i])
	}

	result := 1
	for i := 0; i < n; i++ {
		lcm := i + 1
		count := (r[i] / lcm) - ((l[i] - 1) / lcm)
		result = (result * count) % MOD
	}

	fmt.Fprintln(out, result)
}
