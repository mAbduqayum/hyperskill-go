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
	var n, k int
	fmt.Fscan(in, &n, &k)

	var m int
	fmt.Fscan(in, &m)

	boxes := make([]int, m)
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		boxes[i] = 1 << a
	}

	dp := make([]int, k+1)
	for i := 0; i < m; i++ {
		for j := k; j >= boxes[i]; j-- {
			dp[j] = max(dp[j], dp[j-boxes[i]]+1)
		}
	}

	var trips int
	for i := 0; i <= k; i++ {
		trips = max(trips, dp[i])
	}

	fmt.Fprintln(out, (trips+n-1)/n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
