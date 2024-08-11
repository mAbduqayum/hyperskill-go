package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type Point struct {
	x, y int
}

func main() {
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solveTestCase()
	}
}

func solveTestCase() {
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	resources := make([][]Point, k)
	for i := 0; i < k; i++ {
		var count int
		fmt.Fscan(in, &count)
		resources[i] = make([]Point, count)
		for j := 0; j < count; j++ {
			fmt.Fscan(in, &resources[i][j].x, &resources[i][j].y)
		}
	}

	minArea := n * m
	for _, start := range resources[0] {
		minArea = min(minArea, findMinArea(resources, start, n, m))
	}

	fmt.Fprintln(out, minArea)
}

func findMinArea(resources [][]Point, start Point, n, m int) int {
	xCoords := []int{start.x}
	yCoords := []int{start.y}

	for i := 1; i < len(resources); i++ {
		minDist := n + m
		var bestPoint Point
		for _, point := range resources[i] {
			dist := abs(point.x-start.x) + abs(point.y-start.y)
			if dist < minDist {
				minDist = dist
				bestPoint = point
			}
		}
		xCoords = append(xCoords, bestPoint.x)
		yCoords = append(yCoords, bestPoint.y)
	}

	sort.Ints(xCoords)
	sort.Ints(yCoords)

	return (xCoords[len(xCoords)-1] - xCoords[0] + 1) * (yCoords[len(yCoords)-1] - yCoords[0] + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
