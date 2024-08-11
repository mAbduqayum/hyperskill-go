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
	var n, m int
	fmt.Fscan(in, &n, &m)
	var grid = make([][]int, n)
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		grid[i] = make([]int, m)
		for j, char := range s {
			grid[i][j] = int(char - '0')
		}
	}
	findIJ(n, m, grid)
}

func findIJ(n int, m int, grid [][]int) {
	rowSummary := make([]int, n)
	for i := 0; i < n; i++ {
		rowSummary[i] = smallest(grid[i]...)
	}
	colSummary := make([]int, m)
	for j := 0; j < m; j++ {
		colSummary[j] = 5
		for i := 0; i < n; i++ {
			colSummary[j] = smallest(colSummary[j], grid[i][j])
		}
	}
	worst := gridSmallest(rowSummary, colSummary)
	smallestRow, rs := smallestIndexes(rowSummary...)
	smallestCol, cs := smallestIndexes(colSummary...)
	r := rs[0]
	c := cs[0]

	x, y := -1, -1
	if len(rs) == 1 && len(cs) == 1 {
		//// strike smallestRow against one of smallestCol and check if it's better
		//newColSummary := append(colSummary[:c], colSummary[c+1:]...)
		//for j := 0; j < m; j++ {
		//	w := gridSmallestWithoutIJ(newColSummary, colSummary, r, j)
		//	if w > worst {
		//		worst = w
		//		x, y = smallestRow, j
		//	}
		//}
		//
		//// strike smallestCol against one of smallestRow and check if it's better
		//newRowSummary := append(rowSummary[:r], rowSummary[r+1:]...)
		//for i := 0; i < n; i++ {
		//	w := gridSmallestWithoutIJ(newRowSummary, colSummary, i, c)
		//	if w > worst {
		//		worst = w
		//		x, y = i, smallestCol
		//	}
		//}
	} else if len(rs) == 1 {

		var newColSummary []int
		for j := 0; j < m; j++ {
			if colSummary[j] == smallestCol {
				continue
			}
			newColSummary = append(newColSummary, colSummary[j])
		}
		newRowSummary := append(rowSummary[:r], rowSummary[r+1:]...)
		for i := 0; i < n; i++ {
			wRow := smallestStrikeI(newRowSummary, i)
			wCol := smallest(newColSummary...)
			w := smallest(wRow, wCol)
			if w > worst {
				worst = w
				x, y = i, c
			}
		}
	} else if len(cs) == 1 {
	}

	fmt.Fprintln(out, smallestRow)

	fmt.Fprintln(out, x+1, y+1)
}

func smallest(nums ...int) int {
	rez := nums[0]
	for _, num := range nums {
		if num < rez {
			rez = num
		}
	}
	return rez
}

func smallestIndexes(nums ...int) (int, []int) {
	rez := nums[0]
	indexes := []int{0}
	for i, num := range nums {
		if num < rez {
			rez = num
			indexes = []int{i}
		} else if num == rez {
			indexes = append(indexes, i)
		}
	}
	return rez, indexes
}

func smallestIndex(nums ...int) int {
	rez := 0
	for i, num := range nums {
		if num < nums[rez] {
			rez = i
		}
	}
	return rez
}

func smallestStrikeI(summary []int, index int) int {
	rez := 5
	for i, num := range summary {
		if i != index && num < rez {
			rez = num
		}
	}
	return rez
}

func gridSmallest(rowSummary, colSummary []int) int {
	rowSmallest := smallest(rowSummary...)
	colSmallest := smallest(colSummary...)
	return smallest(rowSmallest, colSmallest)
}

//func gridSmallestWithoutIJ(rowSummary, colSummary []int, i, j int) int {
//	rowSmallest := smallestStrikeI(rowSummary, j)
//	colSmallest := smallestStrikeI(colSummary, i)
//	return smallest(rowSmallest, colSmallest)
//}
