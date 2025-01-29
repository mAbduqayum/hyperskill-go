package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type Box struct {
	Id       string
	X1, Y1   int
	X2, Y2   int
	Area     int
	Parent   *Box
	Children []*Box
}

func main() {
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	results := make([]map[string]interface{}, t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		matrix := readMatrix(n)
		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, m)
		}
		boxes := parseBoxes(matrix, visited)
		findParents(boxes)
		buildTree(boxes)
		rootBoxes := getRootBoxes(boxes)
		result := generateJSON(rootBoxes)
		results[i] = result
	}

	jsonData, _ := json.MarshalIndent(results, "", "  ")
	fmt.Fprintln(out, string(jsonData))
}

func readMatrix(n int) []string {
	matrix := make([]string, n)
	for i := 0; i < n; i++ {
		var line string
		fmt.Fscan(in, &line)
		matrix[i] = line
	}
	return matrix
}

func parseBoxes(matrix []string, visited [][]bool) []*Box {
	var boxes []*Box
	n := len(matrix)
	if n == 0 {
		return boxes
	}
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] || matrix[i][j] != '+' {
				continue
			}

			box := parseBox(matrix, i, j, m)
			if box != nil {
				markVisited(box, visited)
				boxes = append(boxes, box)
			}
		}
	}
	return boxes
}

func parseBox(matrix []string, i, j, maxWidth int) *Box {
	jEnd := -1
	for k := j + 1; k < maxWidth; k++ {
		if matrix[i][k] == '+' {
			jEnd = k
			break
		} else if matrix[i][k] != '-' {
			return nil
		}
	}
	if jEnd == -1 {
		return nil
	}

	iEnd := -1
	for l := i + 1; l < len(matrix); l++ {
		if matrix[l][j] == '+' {
			iEnd = l
			break
		} else if matrix[l][j] != '|' {
			return nil
		}
	}
	if iEnd == -1 {
		return nil
	}

	if len(matrix[iEnd]) < jEnd || matrix[iEnd][jEnd] != '+' {
		return nil
	}
	for k := j + 1; k < jEnd; k++ {
		if len(matrix[iEnd]) <= k || matrix[iEnd][k] != '-' {
			return nil
		}
	}

	for l := i + 1; l < iEnd; l++ {
		if len(matrix[l]) <= jEnd || matrix[l][jEnd] != '|' {
			return nil
		}
	}

	identifier := ""
	if i+1 < len(matrix) {
		row := matrix[i+1]
		for k := j + 1; k < len(row) && len(identifier) < 3; k++ {
			c := row[k]
			if isAlphanumeric(c) {
				identifier += string(c)
			} else {
				break
			}
		}
	}

	width := jEnd - j + 1
	height := iEnd - i + 1
	area := (width - 2) * (height - 2)

	return &Box{
		Id:   identifier,
		X1:   i,
		Y1:   j,
		X2:   iEnd,
		Y2:   jEnd,
		Area: area,
	}
}

func isAlphanumeric(c byte) bool {
	return (c >= '0' && c <= '9') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z')
}

func markVisited(box *Box, visited [][]bool) {
	for y := box.Y1; y <= box.Y2; y++ {
		visited[box.X1][y] = true
		visited[box.X2][y] = true
	}
	for x := box.X1; x <= box.X2; x++ {
		visited[x][box.Y1] = true
		visited[x][box.Y2] = true
	}
}

func findParents(boxes []*Box) {
	for _, child := range boxes {
		var parent *Box
		minArea := math.MaxInt64
		for _, candidate := range boxes {
			if candidate == child {
				continue
			}
			if child.X1 > candidate.X1 && child.X2 < candidate.X2 &&
				child.Y1 > candidate.Y1 && child.Y2 < candidate.Y2 {
				if candidate.Area < minArea {
					parent = candidate
					minArea = candidate.Area
				}
			}
		}
		child.Parent = parent
	}
}

func buildTree(boxes []*Box) {
	for _, box := range boxes {
		if box.Parent != nil {
			box.Parent.Children = append(box.Parent.Children, box)
		}
	}
}

func getRootBoxes(boxes []*Box) []*Box {
	roots := []*Box{}
	for _, box := range boxes {
		if box.Parent == nil {
			roots = append(roots, box)
		}
	}
	return roots
}

func generateJSON(rootBoxes []*Box) map[string]interface{} {
	result := make(map[string]interface{})
	for _, box := range rootBoxes {
		if len(box.Children) == 0 {
			result[box.Id] = box.Area
		} else {
			result[box.Id] = generateChildren(box)
		}
	}
	return result
}

func generateChildren(box *Box) interface{} {
	if len(box.Children) == 0 {
		return box.Area
	}
	childrenMap := make(map[string]interface{})
	for _, child := range box.Children {
		childrenMap[child.Id] = generateChildren(child)
	}
	return childrenMap
}
