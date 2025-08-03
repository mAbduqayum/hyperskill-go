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

type Hexagon struct {
	grid          [][]byte
	width, height int
}

func NewHexagon(width, height int) *Hexagon {
	totalHeight := 2*height + 1
	totalWidth := width + 2*height

	grid := make([][]byte, totalHeight)
	for i := range grid {
		grid[i] = make([]byte, totalWidth)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	// Top line
	for i := 0; i < width; i++ {
		grid[0][height+i] = '_'
	}

	// Upper half
	for row := 1; row <= height; row++ {
		sideIndex := row - 1
		leftPos := height - 1 - sideIndex
		rightPos := leftPos + 1 + width + 2*sideIndex

		grid[row][leftPos] = '/'
		grid[row][rightPos] = '\\'
	}

	// Lower half
	for row := height + 1; row < 2*height; row++ {
		sideIndex := row - height - 1
		leftPos := sideIndex
		rightPos := leftPos + 1 + width + 2*(height-1-sideIndex)

		grid[row][leftPos] = '\\'
		grid[row][rightPos] = '/'
	}

	// Bottom line
	bottomRow := 2 * height
	for i := 0; i < width; i++ {
		grid[bottomRow][height-1+1+i] = '_'
	}
	grid[bottomRow][height-1] = '\\'
	grid[bottomRow][height+width] = '/'

	return &Hexagon{
		grid:   grid,
		width:  width,
		height: height,
	}
}

func (h Hexagon) totalHeight() int {
	return 2*h.height + 1
}

func (h Hexagon) totalWidth() int {
	return h.width + 2*h.height
}

type Board struct {
	grid [][]byte
	m, n int
}

func NewBoard(m, n int) Board {
	grid := make([][]byte, n)
	for i := range grid {
		grid[i] = make([]byte, m)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}
	return Board{
		grid: grid,
		m:    m,
		n:    n,
	}
}

func (b Board) canPlaceHexagon(h *Hexagon, x, y int) bool {
	return x+h.totalWidth() <= b.m && y+h.totalHeight() <= b.n
}

func printBoard(b Board) {
	m := len(b.grid[0])
	n := len(b.grid)

	horizontalBoundary := "+" + strings.Repeat("-", m) + "+\n"
	fmt.Fprint(out, horizontalBoundary)

	for i := 0; i < n; i++ {
		fmt.Fprint(out, "|")
		fmt.Fprint(out, string(b.grid[i]))
		fmt.Fprintln(out, "|")
	}

	fmt.Fprint(out, horizontalBoundary)
}

func (b Board) placeHexagon(h *Hexagon, x, y int) {
	for i := 0; i < h.totalHeight(); i++ {
		for j := 0; j < h.totalWidth(); j++ {
			if h.grid[i][j] != ' ' {
				b.grid[y+i][x+j] = h.grid[i][j]
			}
		}
	}
}

func main() {
	defer out.Flush()

	var m, n, width, height, k int
	fmt.Fscan(in, &m, &n, &width, &height, &k)

	b := NewBoard(m, n)
	h := NewHexagon(width, height)
	x, y := 0, width
	for i := 0; i < k; i++ {
		if b.canPlaceHexagon(h, x, y) {
			b.placeHexagon(h, x, y)
			x += h.totalHeight()
		} else {
			b.placeHexagon(h, x, y)
		}
	}

	printBoard(b)
}
