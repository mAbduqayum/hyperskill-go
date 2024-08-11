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
	var input string
	fmt.Fscan(in, &input)
	terminal(input)
}

func terminal(input string) {
	x, y := 0, 0
	lines := make([][]rune, 1) // Start with one empty input

	for _, char := range input {
		switch char {
		case 'L':
			if x > 0 {
				x--
			}
		case 'R':
			if x < len(lines[y]) {
				x++
			}
		case 'U':
			if y > 0 {
				y--
				if x > len(lines[y]) {
					x = len(lines[y])
				}
			}
		case 'D':
			if y < len(lines)-1 {
				y++
				if x > len(lines[y]) {
					x = len(lines[y])
				}
			}
		case 'B':
			x = 0
		case 'E':
			x = len(lines[y])
		case 'N':
			nextLine := append([]rune{}, lines[y][x:]...)
			lines[y] = lines[y][:x]
			lines = append(lines[:y+1], append([][]rune{nextLine}, lines[y+1:]...)...)
			y++
			x = 0
		default:
			if x == len(lines[y]) {
				lines[y] = append(lines[y], char)
			} else {
				lines[y] = append(lines[y][:x], append([]rune{char}, lines[y][x:]...)...)
			}
			x++
		}
	}

	for _, line := range lines {
		fmt.Fprintln(out, string(line))
	}
	fmt.Fprintln(out, "-")
}
