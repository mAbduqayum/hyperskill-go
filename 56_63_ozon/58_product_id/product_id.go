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

	timeToIds := make([]int, n+1)
	idToProduct := make(map[int]string)
	lastId := 0

	for currentTime := 1; currentTime <= n; currentTime++ {
		var request string
		fmt.Fscan(in, &request)

		if request == "CHANGE" {
			var name string
			var id int
			fmt.Fscan(in, &name, &id)

			lastId = id
			timeToIds[currentTime] = id
			idToProduct[id] = name
		} else if request == "GET" {
			var id, time int
			fmt.Fscan(in, &id, &time)
			timeToIds[currentTime] = lastId
			result := "404"
			if timeToIds[time] != 0 {
				result = idToProduct[timeToIds[time]]
			}
			fmt.Fprintln(out, result)
		}
	}
}
