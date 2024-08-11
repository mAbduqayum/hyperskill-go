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

type Tree struct {
	nodes    map[int]bool
	children map[int]bool
}

func NewTree() *Tree {
	return &Tree{
		nodes:    make(map[int]bool),
		children: make(map[int]bool),
	}
}

func (t *Tree) AddNode(node int) {
	t.nodes[node] = true
}

func (t *Tree) AddChild(child int) {
	t.children[child] = true
}

func (t *Tree) FindRoot() int {
	for node := range t.nodes {
		if !t.children[node] {
			return node
		}
	}
	return -1
}

func main() {
	defer out.Flush()

	var testCases int
	fmt.Fscan(in, &testCases)

	for i := 0; i < testCases; i++ {
		solveTestCase()
	}
}

func solveTestCase() {
	var m int
	fmt.Fscan(in, &m)

	tree := NewTree()

	for i := 0; i < m; {
		var node, childCount int
		fmt.Fscan(in, &node, &childCount)
		tree.AddNode(node)
		i += 2

		for j := 0; j < childCount; j++ {
			var child int
			fmt.Fscan(in, &child)
			tree.AddChild(child)
			i++
		}
	}

	root := tree.FindRoot()
	fmt.Fprintln(out, root)
}
