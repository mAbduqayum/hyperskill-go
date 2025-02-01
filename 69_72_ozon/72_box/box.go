package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type point struct{ x, y int }

type box struct {
	children []box
	id       string
	start    point
	end      point
}

func newBox(id string, start point, end point) box {
	return box{
		id:    id,
		start: start,
		end:   end,
	}
}

func (b box) width() int {
	return b.end.y - b.start.y - 1
}

func (b box) height() int {
	return b.end.x - b.start.x - 1
}

func (b box) area() int {
	return b.width() * b.height()
}

func (b box) toMap() map[string]interface{} {
	if len(b.children) == 0 {
		return map[string]interface{}{
			b.id: b.area(),
		}
	}

	childrenMap := make(map[string]interface{})
	for _, child := range b.children {
		for k, v := range child.toMap() {
			childrenMap[k] = v
		}
	}

	return map[string]interface{}{
		b.id: childrenMap,
	}
}

type logistic struct {
	board []string
	seen  [][]bool
	boxes []box
	m, n  int
}

func newLogistic(board []string) *logistic {
	m, n := len(board), len(board[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	return &logistic{board, seen, nil, m, n}
}

func (l *logistic) parse() {
	start := point{0, 0}
	end := point{l.m, l.n}
	l.boxes = l.getBoxes(start, end)
}

func (l *logistic) getId(x, y int) string {
	id := strings.Builder{}
	for j := y; j < l.n; j++ {
		if !isAlphanumeric(l.board[x][j]) {
			break
		}
		id.WriteByte(l.board[x][j])
	}
	return id.String()
}

func (l *logistic) getBoxes(start, end point) []box {
	var boxes []box
	for i := start.x; i < end.x; i++ {
		for j := start.y; j < end.y; j++ {
			if l.seen[i][j] || l.board[i][j] != '+' {
				continue
			}
			l.seen[i][j] = true

			b := l.getBox(i, j)
			l.markSeen(b)
			s := point{b.start.x + 1, b.start.y + 1}
			e := point{b.end.x - 1, b.end.y - 1}
			b.children = l.getBoxes(s, e)
			boxes = append(boxes, b)
		}
	}
	return boxes
}

func (l *logistic) getBox(x, y int) box {
	start := point{x, y}
	id := l.getId(x+1, y+1)
	endY := y + 1
	for endY < l.n && l.board[x][endY] == '-' {
		endY++
	}
	endX := x + 1
	for endX < l.m && l.board[endX][y] == '|' {
		endX++
	}
	end := point{endX, endY}
	return newBox(id, start, end)
}

func (l *logistic) markSeen(b box) {
	l.seen[b.start.x][b.start.y] = true
	l.seen[b.start.x][b.end.y] = true
	l.seen[b.end.x][b.start.y] = true
	l.seen[b.end.x][b.end.y] = true
}

func main() {
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	results := make([]interface{}, t)
	for i := 0; i < t; i++ {
		results[i] = solveTestCase()
	}

	jsonData, _ := json.MarshalIndent(results, "", "  ")
	fmt.Fprintln(out, string(jsonData))
}

func solveTestCase() interface{} {
	var n, m int
	fmt.Fscan(in, &n, &m)

	board := make([]string, n)
	for i := range n {
		fmt.Fscan(in, &board[i])
	}

	l := newLogistic(board)
	l.parse()

	result := make(map[string]interface{})
	for _, b := range l.boxes {
		for k, v := range b.toMap() {
			result[k] = v
		}
	}

	return result
}

func isAlphanumeric(c byte) bool {
	return c >= '0' && c <= '9' ||
		c >= 'A' && c <= 'Z' ||
		c >= 'a' && c <= 'z'
}
