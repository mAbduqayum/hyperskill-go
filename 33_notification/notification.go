package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	processTest()
}

func processTest() {
	var u, n int
	fmt.Fscan(in, &u, &n)

	var res strings.Builder
	users := make([]int, u)
	var method, id, lastId, lastGlobalId int
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &method, &id)
		if method == 1 {
			lastId++
			if id == 0 {
				lastGlobalId = lastId
			} else {
				users[id-1] = lastId
			}
		} else {
			newId := max(users[id-1], lastGlobalId)
			res.WriteString(strconv.Itoa(newId))
			res.WriteByte('\n')
		}
	}
	fmt.Fprint(out, res.String())
}
