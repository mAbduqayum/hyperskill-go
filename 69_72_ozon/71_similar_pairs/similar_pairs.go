package main

import (
	"bufio"
	"fmt"
	"os"
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

	evenCounts := make(map[string]int)
	oddCounts := make(map[string]int)
	bothCounts := make(map[string]int)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)

		evenKey, oddKey := getKeys(s)

		oddCounts[oddKey]++
		if len(s) == 1 {
			continue
		}
		evenCounts[evenKey]++
		bothCounts[s]++
	}

	evenPairs := countPairs(evenCounts)
	oddPairs := countPairs(oddCounts)
	bothPairs := countPairs(bothCounts)

	total := evenPairs + oddPairs - bothPairs
	fmt.Fprintln(out, total)
}

func getKeys(s string) (string, string) {
	var evenBuilder, oddBuilder strings.Builder
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			oddBuilder.WriteByte(s[i])
		} else {
			evenBuilder.WriteByte(s[i])
		}
	}
	return evenBuilder.String(), oddBuilder.String()
}

func countPairs(m map[string]int) int {
	pairs := 0
	for _, cnt := range m {
		pairs += cnt * (cnt - 1) / 2
	}
	return pairs
}
