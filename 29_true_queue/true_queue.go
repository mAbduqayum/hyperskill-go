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
}

func (dll *DoublyLinkedList[T]) DeleteNode(node *Node[T]) {
	if node == nil {
		return
	}

	// If the node to delete is the head, update the head pointer
	if node == dll.Head {
		dll.Head = node.Next
		if dll.Head != nil { // If the list is not empty after deletion
			dll.Head.Prev = nil
		} else {
			// If the list becomes empty, also update the tail
			dll.Tail = nil
		}
	} else {
		// Update the previous node's next pointer
		node.Prev.Next = node.Next
	}

	// If the node to delete is the tail, update the tail pointer
	if node == dll.Tail {
		dll.Tail = node.Prev
		if dll.Tail != nil { // If the list is not empty after deletion
			dll.Tail.Next = nil
		} else {
			// If the list becomes empty, also update the head
			dll.Head = nil
		}
	} else if node.Next != nil {
		// Update the next node's previous pointer
		node.Next.Prev = node.Prev
	}

	// Clear the pointers of the node to help with garbage collection
	node.Next = nil
	node.Prev = nil
}

func (dll *DoublyLinkedList[T]) Iterate(f func(value T)) {
	current := dll.Head
	for current != nil {
		f(current.Value)
		current = current.Next
	}
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
	dll := NewDoublyLinkedList(s)
	rez := isValidSequence(dll)
	if rez {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

func isValidSequence(dll *DoublyLinkedList[rune]) bool {
	if dll.Head.Value == 'Z' {
		return false
	}
	if dll.Tail.Value == 'X' {
		return false
	}
	for dll.Head.Value == 'Y' {
		current := dll.Head
		for current != nil {
			if current.Value == 'Z' {
				dll.DeleteNode(current)
				break
			}
			current = current.Next
		}
	}
	for dll.Tail.Value == 'Y' {
		current := dll.Tail
		for current != nil {
			if current.Value == 'X' {
				dll.DeleteNode(current)
				break
			}
			current = current.Prev
		}
	}
	current := dll.Head
	for current != nil {
		fmt.Fprintln(out, string(current.Value))
		current = current.Next
	}
	return true
}
