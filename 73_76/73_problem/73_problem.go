package main

import (
	"bufio"
	"fmt"
	"os"
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
		solveTestCase()
	}
}

func solveTestCase() {
	var w, h int
	fmt.Fscan(in, &w, &h)

	totalRows := 2*h + 1

	for row := 0; row < totalRows; row++ {
		switch {
		case row == 0:
			printTopLine(w, h)
		case row <= h:
			printUpperHalf(w, h, row)
		default:
			printLowerHalf(w, h, row, totalRows)
		}
	}
}

func printTopLine(w, h int) {
	leadingSpaces := strings.Repeat(" ", h)
	underscores := strings.Repeat("_", w)
	fmt.Fprintf(out, "%s%s\n", leadingSpaces, underscores)
}

func printUpperHalf(w, h, row int) {
	sideIndex := row - 1
	leadingSpaces := strings.Repeat(" ", h-1-sideIndex)
	middleSpaces := strings.Repeat(" ", w+2*sideIndex)
	fmt.Fprintf(out, "%s/%s\\\n", leadingSpaces, middleSpaces)
}

func printLowerHalf(w, h, row, totalRows int) {
	sideIndex := row - h - 1
	leadingSpaces := strings.Repeat(" ", sideIndex)

	var middle string
	if row == totalRows-1 {
		middle = strings.Repeat("_", w)
	} else {
		middleWidth := w + 2*(h-1-sideIndex)
		middle = strings.Repeat(" ", middleWidth)
	}

	fmt.Fprintf(out, "%s\\%s/\n", leadingSpaces, middle)
}
