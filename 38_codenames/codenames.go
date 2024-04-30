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
	var n, b, r, f int
	fmt.Fscan(in, &n, &b, &r, &f)

	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &words[i])
	}
	wordsSet := arrToSet(words)
	blackWord := words[f-1]
	blackSubWords := subWordsMap(blackWord)
	blueSubWordsMap := getSubWordsMap(words[:b])
	withoutRight(blueSubWordsMap, blackSubWords)
	withoutRight(blueSubWordsMap, wordsSet)
	redSubWordsMap := getSubWordsMap(words[b : b+r])
	withoutRight(redSubWordsMap, blackSubWords)
	withoutRight(redSubWordsMap, wordsSet)
	subWord, bestBR := bestResult(blueSubWordsMap, redSubWordsMap)
	fmt.Fprintln(out, subWord, bestBR)
}

func bestResult(b map[string]int, r map[string]int) (string, int) {
	subWord := "tkhapjiabb"
	var bestBR int
	for k, v := range b {
		newBestBR := v - r[k]
		if newBestBR > bestBR {
			subWord = k
			bestBR = newBestBR
		}
	}
	return subWord, bestBR
}

func subWordsMap(word string) map[string]bool {
	n := len(word)
	rez := make(map[string]bool, n*(n+1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			rez[word[i:j]] = true
		}
	}
	return rez
}

func arrToSet(words []string) map[string]bool {
	rez := make(map[string]bool, len(words))
	for _, word := range words {
		rez[word] = true
	}
	return rez
}

func getSubWordsMap(words []string) map[string]int {
	rez := make(map[string]int)
	for _, w := range words {
		subW := subWordsMap(w)
		for k := range subW {
			if _, ok := rez[k]; ok {
				rez[k]++
			} else {
				rez[k] = 1
			}
		}
	}
	return rez
}

func withoutRight[K comparable, V1 any, V2 any](mapL map[K]V1, mapR map[K]V2) {
	for k := range mapL {
		if _, ok := mapR[k]; ok {
			delete(mapL, k)
		}
	}
}
