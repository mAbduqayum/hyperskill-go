package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type Node struct {
	key      string
	value    string
	children map[string]*Node
}

func newNode(key string) *Node {
	return &Node{
		key:      key,
		children: make(map[string]*Node),
	}
}

func (n *Node) addChild(key string) *Node {
	child := newNode(key)
	n.children[key] = child
	return child
}

func (n *Node) toINI(path []string, result *[]string) {
	if n.value != "" {
		if len(path) > 0 {
			*result = append(*result, fmt.Sprintf("[%s]", strings.Join(path, ".")))
		}
		*result = append(*result, fmt.Sprintf("%s = %s", n.key, n.value))
	}

	childrenKeys := make([]string, 0, len(n.children))
	for k := range n.children {
		childrenKeys = append(childrenKeys, k)
	}

	for _, k := range childrenKeys {
		child := n.children[k]
		newPath := append(path, child.key)
		child.toINI(newPath, result)
	}
}

func parseYAML(lines []string) *Node {
	root := newNode("")
	stack := []*Node{root}
	indentLevel := []int{-1}

	for _, line := range lines {
		indent := 0
		for indent < len(line) && line[indent] == ' ' {
			indent++
		}
		line = strings.TrimSpace(line)

		for indent <= indentLevel[len(indentLevel)-1] {
			stack = stack[:len(stack)-1]
			indentLevel = indentLevel[:len(indentLevel)-1]
		}

		parts := strings.SplitN(line, ": ", 2)
		key := parts[0]

		if len(parts) == 1 {
			node := stack[len(stack)-1].addChild(key)
			stack = append(stack, node)
			indentLevel = append(indentLevel, indent)
		} else {
			stack[len(stack)-1].children[key] = &Node{key: key, value: parts[1]}
		}
	}

	return root
}

func convertYAMLToINI(yaml []string) []string {
	root := parseYAML(yaml)
	var result []string
	for _, child := range root.children {
		child.toINI([]string{}, &result)
		result = append(result, "")
	}
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}

func main() {
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		in.ReadString('\n')

		yaml := make([]string, n)
		for j := 0; j < n; j++ {
			yaml[j], _ = in.ReadString('\n')
			yaml[j] = strings.TrimSpace(yaml[j])
		}

		ini := convertYAMLToINI(yaml)
		for _, line := range ini {
			fmt.Fprintln(out, line)
		}
		fmt.Fprintln(out)
	}
}
