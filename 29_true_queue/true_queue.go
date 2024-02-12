package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Reader
var out *bufio.Writer

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type DoublyLinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func NewDoublyLinkedList(s string) *DoublyLinkedList[rune] {
	dll := &DoublyLinkedList[rune]{}
	for _, runeValue := range s {
		dll.Append(runeValue)
	}
	dll.Size = len(s)
	return dll
}

func (dll *DoublyLinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if dll.Head == nil { // The list is empty
		dll.Head = newNode
		dll.Tail = newNode
	} else { // The list is not empty
		dll.Tail.Next = newNode
		newNode.Prev = dll.Tail
		dll.Tail = newNode
	}
	dll.Size++
}

func (dll *DoublyLinkedList[T]) DeleteNode(node *Node[T]) {
	if node == nil {
		return
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		dll.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		dll.Tail = node.Prev
	}
	dll.Size--
}

func (dll *DoublyLinkedList[T]) Clone() *DoublyLinkedList[T] {
	clone := &DoublyLinkedList[T]{}
	current := dll.Head
	for current != nil {
		clone.Append(current.Value)
		current = current.Next
	}
	return clone
}

func (dll *DoublyLinkedList[T]) Display() {
	current := dll.Head
	for current != nil {
		fmt.Fprintf(out, "%c", current.Value)
		current = current.Next
	}
	fmt.Fprintln(out)
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
	var n int
	fmt.Fscan(in, &n)

	var s string
	fmt.Fscan(in, &s)
	dll := NewDoublyLinkedList[rune](s)
	rez := isValidSequence(dll)
	if rez {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}

func isValidSequence(dll *DoublyLinkedList[rune]) bool {
	if dll.Head != nil && dll.Head.Value == 'Z' || dll.Tail != nil && dll.Tail.Value == 'X' {
		return false
	}

	if dll.Tail.Value == 'X' {
		return false
	}
	for dll.Head != nil && dll.Head.Value == 'Y' {
		foundZ := false
		current := dll.Head.Next
		for current != nil {
			if current.Value == 'Z' {
				foundZ = true
				dll.DeleteNode(current)
				dll.DeleteNode(dll.Head)
				break
			}
			current = current.Next
		}
		if !foundZ {
			return false
		}
	}

	for dll.Tail != nil && dll.Tail.Value == 'Y' {
		foundX := false
		current := dll.Tail.Prev
		for current != nil {
			if current.Value == 'X' {
				foundX = true
				dll.DeleteNode(current)
				dll.DeleteNode(dll.Tail)
				break
			}
			current = current.Prev
		}
		if !foundX {
			break
		}
	}

	for dll.Size > 4 && dll.Tail.Value == 'Z' {
		curr := dll.Tail.Prev
		foundNonZ := false
		foundZ := false
		var nonZ *Node[rune]
		for curr != nil {
			if curr.Value == 'Z' {
				foundZ = true
			} else if !foundNonZ {
				foundNonZ = true
				nonZ = curr
			}
			if foundZ && foundNonZ {
				dll.DeleteNode(dll.Tail)
				dll.DeleteNode(nonZ)
				break
			}
			curr = curr.Prev
		}
		if !foundZ {
			break
		}
	}

	if dll.Size == 0 {
		return true
	}

	if dll.Size == 2 {
		h := dll.Head.Value
		t := dll.Tail.Value
		ht := string(h) + string(t)
		return ht == "XY" || ht == "XZ" || ht == "YZ"
	}

	if dll.Size == 4 {
		if dll.Head.Value == 'X' && dll.Tail.Value == 'Z' {
			i2 := dll.Head.Next.Value
			i3 := dll.Head.Next.Next.Value
			return string(i2)+string(i3) != "XX"
		}
		return isValidSequence(dll)
	}

	clonedDLL := dll.Clone()
	clonedDLL.DeleteNode(clonedDLL.Tail)
	curr := clonedDLL.Tail
	for curr.Value != 'X' {
		curr = curr.Prev
	}
	clonedDLL.DeleteNode(curr)
	if isValidSequence(clonedDLL) {
		return true
	}

	dll.DeleteNode(dll.Tail)
	curr = dll.Tail
	for curr.Value != 'Y' {
		curr = curr.Prev
	}
	dll.DeleteNode(curr)
	return isValidSequence(dll)
}
