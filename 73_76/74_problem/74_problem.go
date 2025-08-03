package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()

	var m, n, width, height, k int
	fmt.Fscan(in, &m, &n, &width, &height, &k)

	grid := make([][]rune, n+2)
	for i := range grid {
		grid[i] = make([]rune, m+2)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	for j := 0; j < m+2; j++ {
		grid[0][j] = '-'
		grid[n+1][j] = '-'
	}
	grid[0][0] = '+'
	grid[0][m+1] = '+'
	grid[n+1][0] = '+'
	grid[n+1][m+1] = '+'

	for i := 1; i <= n; i++ {
		grid[i][0] = '|'
		grid[i][m+1] = '|'
	}

	placed := 0
	for row := 0; placed < k; row++ {
		for col := 0; placed < k; col++ {
			var x, y int
			if row%2 == 0 {
				x = col*(width+height) + 1
				y = row*height + 1
			} else {
				x = col*(width+height) + height + 1
				y = row*height + 1
			}

			hexWidth := width + 2*height
			hexHeight := 2*height + 1

			if y+hexHeight-1 <= n && x+hexWidth-1 <= m {
				drawHexagon(grid, x, y, width, height)
				placed++
			} else if col == 0 {
				break
			} else {
				break
			}
		}
	}

	for i := 0; i < n+2; i++ {
		for j := 0; j < m+2; j++ {
			fmt.Fprintf(out, "%c", grid[i][j])
		}
		fmt.Fprintln(out)
	}
}

func drawHexagon(grid [][]rune, startX, startY, width, height int) {
	for i := 0; i < width; i++ {
		grid[startY][startX+height+i] = '_'
		grid[startY+2*height][startX+height+i] = '_'
	}

	for i := 0; i < height; i++ {
		grid[startY+1+i][startX+height-1-i] = '/'
		grid[startY+1+i][startX+height+width+i] = '\\'
		grid[startY+height+1+i][startX+i] = '\\'
		grid[startY+height+1+i][startX+height+width+i] = '/'
	}
}
