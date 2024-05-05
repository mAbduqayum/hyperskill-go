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

	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i] > boxes[j]
	})

	var trips int
	var currentLoad int
	skip := make([]bool, m)
	for i, box := range boxes {
		if skip[i] {
			continue
		}
		if currentLoad+box > k {
			for j := i + 1; j < m; j++ {
				if !skip[j] && currentLoad+boxes[j] <= k {
					currentLoad += boxes[j]
					skip[j] = true
				}
			}
			trips++
			currentLoad = box
		} else {
			currentLoad += box
		}
	}
	if currentLoad > 0 {
		trips++
	}

	fmt.Fprintln(out, (trips+n-1)/n)
}
