package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	count := 0
	cache := make(map[int][]int, n)

	for i := 1; i < n-1; i++ {
		leftDiff := a[i] - a[i-1]

		if indexes, exists := cache[leftDiff]; exists {
			count += countGreaterThan(indexes, i)
			continue
		}

		for j := i + 1; j < n-1; j++ {
			rightDiff := a[j+1] - a[j]
			if leftDiff == rightDiff {
				cache[leftDiff] = append(cache[leftDiff], j)
				count++
			}
		}
	}

	fmt.Fprintln(out, count)
}

func countGreaterThan(sortedSlice []int, i int) int {
	index := sort.SearchInts(sortedSlice, i+1)
	return len(sortedSlice) - index
}
