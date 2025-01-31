package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pair struct {
	even, odd string
}

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
	pairCounts := make(map[pair]int)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)

		evenKey, oddKey := getKeys(s)

		evenCounts[evenKey]++
		oddCounts[oddKey]++
		pairCounts[pair{evenKey, oddKey}]++
	}

	delete(evenCounts, "")
	delete(oddCounts, "")

	evenPairs := 0
	for _, cnt := range evenCounts {
		evenPairs += cnt * (cnt - 1) / 2
	}

	oddPairs := 0
	for _, cnt := range oddCounts {
		oddPairs += cnt * (cnt - 1) / 2
	}

	bothPairs := 0
	for p, cnt := range pairCounts {
		if p.even != "" && p.odd != "" {
			bothPairs += cnt * (cnt - 1) / 2
		}
	}

	total := evenPairs + oddPairs - bothPairs
	fmt.Fprintln(out, total)
}

func getKeys(s string) (string, string) {
	var evenBuilder, oddBuilder strings.Builder
	for i := 0; i < len(s); i++ {
		if i%2 == 1 {
			evenBuilder.WriteByte(s[i])
		} else {
			oddBuilder.WriteByte(s[i])
		}
	}
	return evenBuilder.String(), oddBuilder.String()
}
