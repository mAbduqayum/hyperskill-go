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
	var n, t int
	fmt.Fscan(in, &n, &t)
	letters := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &letters[i])
	}
	for i := 0; i < t; i++ {
		processTest(letters)
	}
}

func processTest(letters []string) {
	var pinCode string
	fmt.Fscan(in, &pinCode)
	if isValidPinCode(pinCode, letters) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

func isValidPinCode(pinCode string, letters []string) bool {
	if len(pinCode) != len(letters) {
		return false
	}
	freq := make(map[string]int)
	for _, c := range pinCode {
		freq[strings.ToLower(string(c))]++
	}
	for _, letter := range letters {
		if freq[letter] == 0 {
			return false
		}
		freq[letter]--
	}
	return true
}
