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
	var k int
	fmt.Fscan(in, &k)
	var s string
	fmt.Fscan(in, &s)
	printDocument(k, s)
}

func printDocument(totalPages int, printedStr string) {
	printed := make([]bool, totalPages+2)
	printed[0] = true            // to skip
	printed[totalPages+1] = true // to skip

	bounds := strings.Split(printedStr, ",")
	for _, bound := range bounds {
		if strings.Contains(bound, "-") {
			parts := strings.Split(bound, "-")
			l, _ := strconv.Atoi(parts[0])
			r, _ := strconv.Atoi(parts[1])
			for i := l; i <= r; i++ {
				printed[i] = true
			}
		} else {
			page, _ := strconv.Atoi(bound)
			printed[page] = true
		}
	}

	var rez []string
	for i := 1; i <= totalPages+1; i++ {
		if !printed[i] {
			l := i
			for i <= totalPages+1 && !printed[i] {
				i++
			}
			r := i - 1
			if l == r {
				rez = append(rez, strconv.Itoa(l))
			} else {
				rez = append(rez, fmt.Sprint(l, "-", r))
			}
		}
	}
	fmt.Fprintln(out, strings.Join(rez, ","))
}
