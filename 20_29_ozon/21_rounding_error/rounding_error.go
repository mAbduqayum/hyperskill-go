package main

import (
	"bufio"
	"fmt"
	"math"
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
	var n, p int
	fmt.Fscan(in, &n, &p)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Fprintf(out, "%.2f\n", solve(p, a))
}

func solve(p int, a []int) float64 {
	var lost float64
	for _, ai := range a {
		profit := float64(ai) * float64(p) / 100
		lost += profit - math.Floor(profit)
	}
	return lost
}
