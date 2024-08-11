package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	processTest()
}

func processTest() {
	var n, m int
	fmt.Fscan(in, &n, &m)
	friends := make([]int, n)
	for i := range friends {
		fmt.Fscan(in, &friends[i])
	}
	sf := make([]int, n)
	copy(sf, friends)
	sort.Sort(sort.Reverse(sort.IntSlice(sf)))

	ids := make(map[int]int)
	for i, s := range sf {
		if _, ok := ids[s]; !ok {
			ids[s] = m - i
		}
	}

	res := make([]int, n)
	for i, f := range friends {
		res[i] = ids[f]
		ids[f]--
	}
	for i, r := range res {
		if r <= friends[i] {
			fmt.Fprintln(out, -1)
			return
		}
	}
	var result strings.Builder
	for _, r := range res {
		result.WriteString(strconv.Itoa(r))
		result.WriteByte(' ')
	}
	fmt.Fprintln(out, result.String())
}
