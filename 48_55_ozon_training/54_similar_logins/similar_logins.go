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

	var n int
	fmt.Fscan(in, &n)

	existingLogins := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &existingLogins[i])
	}

	var m int
	fmt.Fscan(in, &m)

	for i := 0; i < m; i++ {
		var newLogin string
		fmt.Fscan(in, &newLogin)
		if hasSimilarLogin(newLogin, existingLogins) {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}

func hasSimilarLogin(newLogin string, existingLogins []string) bool {
	for _, login := range existingLogins {
		if isSimilar(newLogin, login) {
			return true
		}
	}
	return false
}

func isSimilar(a, b string) bool {
	if a == b {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	if diff != 2 {
		return false
	}
	for i := 0; i < len(a)-1; i++ {
		swapped := []byte(a)
		swapped[i], swapped[i+1] = swapped[i+1], swapped[i]
		if string(swapped) == b {
			return true
		}
	}
	return false
}
