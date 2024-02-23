package main

import (
	"bufio"
	"fmt"
	"os"
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
	var n, x, y, xY, xyxz int
	fmt.Fscan(in, &n)
	var chars string
	fmt.Fscan(in, &chars)
	flag := true
	for _, char := range chars {
		switch char {
		case 'X':
			x++
		case 'Y':
			if x > 0 {
				x--
				xY++
			} else if xyxz > 0 {
				xyxz--
				xY++
				x++
			} else {
				y++
			}
		default:
			if y > 0 {
				y--
			} else if x > 0 {
				x--
				if xY > 0 {
					xyxz++
					xY--
				}
			} else if xY > 0 {
				xY--
				x++
			} else if xyxz > 0 {
				xyxz--
				x++
			} else {
				flag = false
			}
		}
	}
	if flag && x == 0 && y == 0 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
