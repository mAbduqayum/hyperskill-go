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

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solveTestCase()
	}
}

func solveTestCase() {
	var n, m int
	fmt.Fscan(in, &n, &m)

	board := make([][]byte, n)
	var row string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &row)
		board[i] = []byte(row)
	}
	solve(n, m, board)
}

type Point struct {
	x, y int
}

func solve(n int, m int, board [][]byte) {
	//origin := Point{0, 0}
	//end := Point{n, m}
	//aCoord := getCoord(n, m, 'A', board)
	//bCoord := getCoord(n, m, 'B', board)
	// logic here
	beautify(board)
}

func getCoord(n int, m int, wanted byte, board [][]byte) Point {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == wanted {
				return Point{i, j}
			}
		}
	}
	return Point{-1, -1}
}

func beautify(board [][]byte) {
	for i := 0; i < len(board); i++ {
		out.Write(board[i])
		out.WriteByte('\n')
	}
}
