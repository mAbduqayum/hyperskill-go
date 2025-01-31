package main4

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type StringPattern struct {
	even []byte
	odd  []byte
}

var patternCache map[string]StringPattern

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

	patternCache = make(map[string]StringPattern, n)

	strings := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &strings[i])
	}

	slices.SortFunc(strings, func(a, b string) int {
		return len(a) - len(b)
	})

	similarPairs := 0
	for i := 0; i < len(strings)-1; i++ {
		for j := i + 1; j < len(strings); j++ {
			pattern1 := getPattern(strings[i])
			if len(strings[i])-len(strings[j]) < -1 {
				break
			}
			pattern2 := getPattern(strings[j])
			if areSimilar(pattern1, pattern2) {
				similarPairs++
			}
		}
	}
	fmt.Fprintln(out, similarPairs)
}

func getPattern(s string) StringPattern {
	if pattern, exists := patternCache[s]; exists {
		return pattern
	}
	odd := make([]byte, 0, len(s)/2+1)
	even := make([]byte, 0, len(s)/2)
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			even = append(even, s[i])
		} else {
			odd = append(odd, s[i])
		}
	}
	pattern := StringPattern{even, odd}
	patternCache[s] = pattern
	return pattern
}

func areSimilar(pattern1, pattern2 StringPattern) bool {
	if len(pattern1.odd) == 0 || len(pattern2.odd) == 0 {
		return slices.Equal(pattern1.even, pattern2.even)
	}
	return slices.Equal(pattern1.even, pattern2.even) || slices.Equal(pattern1.odd, pattern2.odd)
}
