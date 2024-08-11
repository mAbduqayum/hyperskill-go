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
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		processTest()
	}
}

func processTest() {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	memo := make(map[int]int)
	s := make([]int, n)
	copy(s, a)
	sort.Ints(s)
	id, lastId := 1, 1
	memo[s[0]] = lastId
	for i := 1; i < n; i++ {
		id++
		if s[i]-s[i-1] > 1 {
			lastId = id
		}
		memo[s[i]] = lastId
	}
	var rez strings.Builder
	for _, v := range a {
		rez.WriteString(strconv.Itoa(memo[v]))
		rez.WriteByte(' ')
	}
	fmt.Fprintln(out, rez.String())
}
