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
	var k, n, m int
	fmt.Fscan(in, &k, &n, &m)

	board := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &board[i])
	}

	if hasWinner(board, k) {
		fmt.Fprintln(out, "NO")
		return
	}

	if canWinWithOneMove(board, k) {
		fmt.Fprintln(out, "YES")
		return
	}
	fmt.Fprintln(out, "NO")
}

func hasWinner(board []string, k int) bool {
	n, m := len(board), len(board[0])

	// Check rows and columns
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] != '.' {
				if checkSequence(board, i, j, 0, 1, k) ||
					checkSequence(board, i, j, 1, 0, k) ||
					checkSequence(board, i, j, 1, 1, k) ||
					checkSequence(board, i, j, 1, -1, k) {
					return true
				}
			}
		}
	}

	return false
}

func checkSequence(board []string, i, j, di, dj, k int) bool {
	n, m := len(board), len(board[0])
	symbol := board[i][j]
	count := 0

	for i >= 0 && i < n && j >= 0 && j < m && board[i][j] == symbol {
		count++
		if count == k {
			return true
		}
		i += di
		j += dj
	}

	return false
}

func canWinWithOneMove(board []string, k int) bool {
	n, m := len(board), len(board[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == '.' {
				if checkWinningMove(board, i, j, k) {
					return true
				}
			}
		}
	}

	return false
}

func checkWinningMove(board []string, i, j, k int) bool {
	directions := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}

	for _, dir := range directions {
		count := 1
		ni, nj := i+dir[0], j+dir[1]

		for count < k && isValid(board, ni, nj) && board[ni][nj] == 'X' {
			count++
			ni += dir[0]
			nj += dir[1]
		}

		ni, nj = i-dir[0], j-dir[1]
		for count < k && isValid(board, ni, nj) && board[ni][nj] == 'X' {
			count++
			ni -= dir[0]
			nj -= dir[1]
		}

		if count == k {
			return true
		}
	}

	return false
}

func isValid(board []string, i, j int) bool {
	return i >= 0 && i < len(board) && j >= 0 && j < len(board[0])
}
