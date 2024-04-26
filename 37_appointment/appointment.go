package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type slot struct {
	order int
	val   int
	sign  byte
}

func processTest() {
	var n, m int
	fmt.Fscan(in, &n, &m)
	slots := make([]int, m)
	sortedslots := make([]*slot, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &slots[i])
		sortedslots[i] = &slot{i, slots[i], '0'}
	}

	sort.Slice(sortedslots, func(i, j int) bool { return sortedslots[i].val < sortedslots[j].val })

	for i, s := range sortedslots {
		if i == 0 {
			if s.val == 0 {
				s.val++
				s.sign = '+'
			} else if s.val > 1 {
				s.val--
				s.sign = '-'
			}
			continue
		}
		gap := s.val - sortedslots[i-1].val
		if gap == 0 {
			s.val++
			s.sign = '+'
			continue
		}
		if gap > 1 {
			s.val--
			s.sign = '-'
			continue
		}
	}
	sort.Slice(sortedslots, func(i, j int) bool { return sortedslots[i].order < sortedslots[j].order })
	var rez strings.Builder
	for _, s := range sortedslots {
		rez.WriteByte(s.sign)
	}
	fmt.Println(rez.String())
}
