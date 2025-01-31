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

	evenMap := make(map[string]int, n)
	oddMap := make(map[string]int, n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)

		even := evenPattern(s)
		odd := oddPattern(s)
		evenMap[even]++
		if len(odd) == 0 {
			continue
		}
		oddMap[odd]++
	}

	pairs := 0

	for _, count := range evenMap {
		if count >= 2 {
			pairs += comb(count, 2)
		}
	}

	for _, count := range oddMap {
		if count >= 2 {
			pairs += comb(count, 2)
		}
	}

	for key := range evenMap {
		if oddMap[key] > 0 {
			pairs -= evenMap[key] * oddMap[key]
		}
	}

	fmt.Fprintln(out, pairs)
}

func evenPattern(s string) string {
	var evenChars []byte
	for i := 0; i < len(s); i += 2 {
		evenChars = append(evenChars, s[i])
	}
	return string(evenChars)
}

func oddPattern(s string) string {
	var oddChars []byte
	for i := 1; i < len(s); i += 2 {
		oddChars = append(oddChars, s[i])
	}
	return string(oddChars)
}

var combCache = make(map[[2]int]int)

func comb(m, n int) int {
	key := [2]int{m, n}
	if val, ok := combCache[key]; ok {
		return val
	}
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
	combCache[key] = result
	return result
}
