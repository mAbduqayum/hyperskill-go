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

	printTopLine(w, h)
	for row := 1; row <= h; row++ {
		printUpperHalf(w, h, row)
	}
	for row := h + 1; row < 2*h; row++ {
		printLowerHalf(w, h, row)
	}
	printBottomLine(w, h)
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

func printLowerHalf(w, h, row int) {
	sideIndex := row - h - 1
	leadingSpaces := strings.Repeat(" ", sideIndex)
	middleWidth := w + 2*(h-1-sideIndex)
	middleSpaces := strings.Repeat(" ", middleWidth)
	fmt.Fprintf(out, "%s\\%s/\n", leadingSpaces, middleSpaces)
}

func printBottomLine(w, h int) {
	leadingSpaces := strings.Repeat(" ", h-1)
	underscores := strings.Repeat("_", w)
	fmt.Fprintf(out, "%s\\%s/\n", leadingSpaces, underscores)
}
