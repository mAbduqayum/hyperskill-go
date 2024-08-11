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
	var n, k, m int
	fmt.Fscan(in, &n, &k, &m)

	boxCounts := make([]int, 30)
	for i := 0; i < m; i++ {
		var a int
		fmt.Fscan(in, &a)
		boxCounts[a]++
	}

	trips := 0
	trucks := n

	for i := 29; i >= 0; i-- {
		boxWeight := 1 << i
		for boxCounts[i] > 0 {
			if trucks == 0 {
				trips++
				trucks = n
			}
			spaceLeft := k
			for spaceLeft >= boxWeight && boxCounts[i] > 0 {
				spaceLeft -= boxWeight
				boxCounts[i]--
			}
			for j := i - 1; j >= 0; j-- {
				smallerBoxWeight := 1 << j
				for spaceLeft >= smallerBoxWeight && boxCounts[j] > 0 {
					spaceLeft -= smallerBoxWeight
					boxCounts[j]--
				}
			}
			trucks--
		}
	}

	if trucks < n {
		trips++
	}

	fmt.Fprintln(out, trips)
}
