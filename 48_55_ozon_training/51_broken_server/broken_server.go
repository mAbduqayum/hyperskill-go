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
		solveTestCase()
	}
}

func solveTestCase() {
	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	maxLen := 0
	left := 0
	clientCount := make(map[int]int)

	for right := 0; right < n; right++ {
		clientCount[a[right]]++

		for len(clientCount) > 2 {
			clientCount[a[left]]--
			if clientCount[a[left]] == 0 {
				delete(clientCount, a[left])
			}
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	fmt.Fprintln(out, maxLen)
}
