package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in *bufio.Reader
var out *bufio.Writer

type point struct {
	x int
	y int
}

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
	var row, col int
	fmt.Fscan(in, &row, &col)
	var board [][]string
	for i := 0; i < row; i++ {
		var row string
		fmt.Fscan(in, &row)
		board = append(board, strings.Split(row, ""))
	}
	solve(board, row, col)
}

func solve(board [][]string, row int, col int) {
	origin := &point{
		x: 0,
		y: 0,
	}
	destination := &point{
		x: row - 1,
		y: col - 1,
	}
	var pointA *point
	var pointB *point
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == "A" {
				pointA = &point{
					x: i,
					y: j,
				}
			}
			if board[i][j] == "B" {
				pointB = &point{
					x: i,
					y: j,
				}
			}
		}
	}
	if squaredDistance(*origin, *pointA) < squaredDistance(*origin, *pointB) {
		fillPath(board, origin, pointA, "a")
		fillPath(board, destination, pointB, "b")
	} else {
		fillPath(board, origin, pointB, "b")
		fillPath(board, destination, pointA, "a")
	}
	printBoard(board)
}

func squaredDistance(p1, p2 point) int {
	return (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
}

func fillPath(board [][]string, start, end *point, paver string) {
	if end.x%2 == 0 {
		if start.x <= end.x {
			for i := start.x; i < end.x; i++ {
				board[i][start.y] = paver
			}
		} else {
			for i := start.x; i > end.x; i-- {
				board[i][start.y] = paver
			}
		}
		if start.y <= end.y {
			for j := start.y; j < end.y; j++ {
				board[end.x][j] = paver
			}
		} else {
			for j := start.y; j > end.y; j-- {
				board[end.x][j] = paver
			}
		}
	} else {
		if start.y <= end.y {
			for j := start.y; j < end.y; j++ {
				board[start.x][j] = paver
			}
		} else {
			for j := start.y; j > end.y; j-- {
				board[start.x][j] = paver
			}
		}
		if start.x <= end.x {
			for i := start.x; i < end.x; i++ {
				board[i][end.y] = paver
			}
		} else {
			for i := start.x; i > end.x; i-- {
				board[i][end.y] = paver
			}
		}
	}
}

func printBoard(board [][]string) {
	//fmt.Fprintln(out)
	for _, row := range board {
		fmt.Fprintln(out, strings.Join(row, ""))
	}
}
