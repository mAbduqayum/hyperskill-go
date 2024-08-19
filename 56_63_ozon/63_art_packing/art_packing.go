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

type Box struct {
	width, height int
}

type Painting struct {
	width, height int
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
	var n int
	fmt.Fscan(in, &n)

	boxes := make([]Box, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &boxes[i].width, &boxes[i].height)
		if boxes[i].width > boxes[i].height {
			boxes[i].width, boxes[i].height = boxes[i].height, boxes[i].width
		}
	}

	var m int
	fmt.Fscan(in, &m)

	paintings := make([]Painting, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &paintings[i].width, &paintings[i].height)
		if paintings[i].width > paintings[i].height {
			paintings[i].width, paintings[i].height = paintings[i].height, paintings[i].width
		}
	}

	// Sort boxes and paintings by width, then by height
	sort.Slice(boxes, func(i, j int) bool {
		if boxes[i].width == boxes[j].width {
			return boxes[i].height < boxes[j].height
		}
		return boxes[i].width < boxes[j].width
	})

	sort.Slice(paintings, func(i, j int) bool {
		if paintings[i].width == paintings[j].width {
			return paintings[i].height < paintings[j].height
		}
		return paintings[i].width < paintings[j].width
	})

	boxCount := 0
	paintingIndex := 0

	for _, box := range boxes {
		for paintingIndex < m && paintings[paintingIndex].width <= box.width && paintings[paintingIndex].height <= box.height {
			paintingIndex++
		}
		if paintingIndex > 0 {
			boxCount++
			paintingIndex = 0
		}
	}

	if paintingIndex < m {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, boxCount)
	}
}
