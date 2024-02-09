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
	var line string
	fmt.Fscanln(in, &line)
	terminal(line)
}

func terminal(line string) {
	x := 0
	y := 0
	rez := make([][]rune, 0)
	for _, s := range line {
		switch s {
		case 'L':
			if x > 0 {
				x--
			}
			continue
		case 'R':
			if y < len(rez[y]) {
				x++
			}
			continue
		case 'U':
			if y > 0 {
				y--
			}
			continue
		case 'D':
			if y < len(rez)-1 {
				y++
			}
			continue
		case 'H':
			x = 0
			continue
		case 'E':
			x = len(rez[y])
			continue
		case 'N':
			var nextLine []rune
			if x < len(rez[y])-1 {
				rez[y] = rez[y][:x]
				nextLine = rez[y][x:]
			}
			others := rez[y+1]
			rez = append(rez[:y], append(nextLine))
			rez = append(rez, others)
		default:
			rez[y] = append(rez[y], s)
			x++
		}
	}
	fmt.Println(rez)
	fmt.Fprintln(out, "-")
}
