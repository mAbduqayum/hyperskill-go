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
	sortedSlots := make([]*slot, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &slots[i])
		sortedSlots[i] = &slot{i, slots[i], '0'}
	}

	sort.Slice(sortedSlots, func(i, j int) bool { return sortedSlots[i].val < sortedSlots[j].val })

	for i, s := range sortedSlots {
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
		gap := s.val - sortedSlots[i-1].val
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
		if gap < 0 {
			fmt.Fprintln(out, "x")
			return
		}
	}
	var biggest int
	for _, s := range sortedSlots {
		if s.val > biggest {
			biggest = s.val
		}
	}
	if biggest > n || biggest < m {
		fmt.Fprintln(out, "x")
		return
	}
	sort.Slice(sortedSlots, func(i, j int) bool { return sortedSlots[i].order < sortedSlots[j].order })
	var rez strings.Builder
	for _, s := range sortedSlots {
		rez.WriteByte(s.sign)
	}
	fmt.Fprintln(out, rez.String())
}
