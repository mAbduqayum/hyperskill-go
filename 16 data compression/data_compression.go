package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in *bufio.Reader
var out *bufio.Writer

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		processTest()
	}
}

func processTest() {
	var count int
	fmt.Fscan(in, &count)
	data := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Fscan(in, &data[i])
	}
	compress(data)
}

func compress(data []int) {
	compressed := make([]int, 0)
	l := 0
	r := len(data) - 1
	for l <= r {
		start := data[l]
		length := 0
		if l < r && data[l] == data[l+1]-1 {
			for l < r && data[l] == data[l+1]-1 {
				length++
				l++
			}
		} else if l < r && data[l] == data[l+1]+1 {
			for l < r && data[l] == data[l+1]+1 {
				length--
				l++
			}
		}
		l++
		compressed = append(compressed, start, length)
	}
	rezStrings := make([]string, len(compressed))
	for i, num := range compressed {
		rezStrings[i] = strconv.Itoa(num)
	}
	fmt.Fprintln(out, len(compressed))
	fmt.Fprintln(out, strings.Join(rezStrings, " "))
}
