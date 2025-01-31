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
	oddHash := make([]byte, n)
	evenhash := make([]byte, n)

	for i := range n {
		var slogan string
		fmt.Fscan(in, &slogan)
		oddKey := make([]byte, len(slogan)/2+1)
		for j := 0; j < len(oddKey); j++ {
			oddKey[j] = slogan[2*j+1]
		}
		evenKey := make([]byte, len(slogan)/2)
		for j := 0; j < len(evenKey); j++ {
			evenKey[j] = slogan[2*j]
		}
		oddKety
	}
}

func comb(m, n int) int {
	if n == 0 || n == m {
		return 1
	}
	if n > m/2 {
		n = m - n
	}
	result := 1
	for i := 0; i < n; i++ {
		result *= m - i
		result /= i + 1
	}
	return result
}
