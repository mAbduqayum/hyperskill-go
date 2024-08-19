package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
		processTest()
	}
}

type truck struct {
	start    int
	end      int
	capacity int
	// todo: try to remove index
	index int
}

func processTest() {
	var arrivalsCount int
	fmt.Fscan(in, &arrivalsCount)

	arrivals := make([]int, arrivalsCount)
	for i := 0; i < arrivalsCount; i++ {
		fmt.Fscan(in, &arrivals[i])
	}

	var trucksCount int
	fmt.Fscan(in, &trucksCount)

	trucks := make([]truck, trucksCount)
	for i := 0; i < trucksCount; i++ {
		fmt.Fscan(in, &trucks[i].start, &trucks[i].end, &trucks[i].capacity)
		trucks[i].index = i
	}

	sortedArrivals := make([]int, arrivalsCount)
	copy(sortedArrivals, arrivals)
	sort.Ints(sortedArrivals)
	sortedArrivalsMap := make(map[int]int, arrivalsCount)

	sortedTrucks := make([]truck, trucksCount)
	copy(sortedTrucks, trucks)
	sort.Slice(sortedTrucks, func(i, j int) bool {
		if sortedTrucks[i].start == sortedTrucks[j].start {
			return sortedTrucks[i].index < sortedTrucks[j].index
		}
		return sortedTrucks[i].start < sortedTrucks[j].start
	})

	currentTruckIndex := 0
	for i := 0; i < arrivalsCount; i++ {
		arrival := sortedArrivals[i]
		if arrival < sortedTrucks[currentTruckIndex].start {
			continue
		}
		if arrival > sortedTrucks[currentTruckIndex].end {
			continue
		}
		sortedArrivalsMap[arrival] = sortedTrucks[currentTruckIndex].index + 1
		sortedTrucks[currentTruckIndex].capacity--
		if sortedTrucks[currentTruckIndex].capacity == 0 {
			currentTruckIndex++
			if currentTruckIndex >= trucksCount {
				break
			}
		}
	}

	var rez strings.Builder
	for _, arrival := range arrivals {
		if val, ok := sortedArrivalsMap[arrival]; ok {
			rez.WriteString(strconv.Itoa(val))
		} else {
			rez.WriteString(strconv.Itoa(-1))
		}
		rez.WriteByte(' ')
	}
	fmt.Fprintln(out, rez.String())
}
