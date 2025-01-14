package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	var n int
	fmt.Fscan(in, &n)

	items := make([]Item, n)
	for i := 0; i < n; i++ {
		items[i] = Item{index: i + 1, truckIndex: -1}
		fmt.Fscan(in, &items[i].arrival)
	}

	var m int
	fmt.Fscan(in, &m)

	trucks := make([]Truck, m)
	for i := 0; i < m; i++ {
		trucks[i] = Truck{index: i + 1}
		fmt.Fscan(in, &trucks[i].start, &trucks[i].end, &trucks[i].capacity)
	}

	solve(items, trucks)
}

type Item struct {
	index      int
	arrival    int
	truckIndex int
}

type Truck struct {
	index    int
	start    int
	end      int
	capacity int
}

func solve(items []Item, trucks []Truck) {
	slices.SortFunc(items, func(a, b Item) int {
		return a.arrival - b.arrival
	})
	slices.SortFunc(trucks, func(a, b Truck) int {
		if a.start == b.start {
			return a.index - b.index
		}
		return a.start - b.start
	})

	currentItemIndex := 0
	for truckIdx := range trucks {
		truck := &trucks[truckIdx]
		remainingCapacity := truck.capacity

		for currentItemIndex < len(items) && remainingCapacity > 0 {
			if items[currentItemIndex].arrival > truck.end {
				break
			}
			if items[currentItemIndex].arrival >= truck.start {
				items[currentItemIndex].truckIndex = truck.index
				remainingCapacity--
			}
			currentItemIndex++
		}
	}

	slices.SortFunc(items, func(a, b Item) int {
		return a.index - b.index
	})

	for _, item := range items {
		fmt.Fprint(out, item.truckIndex, " ")
	}
	fmt.Fprintln(out)
}
