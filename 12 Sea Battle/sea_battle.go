package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	t, _ := strconv.Atoi(input())
	result := make([]bool, t)
	for i := 0; i < t; i++ {
		ships := readShips()
		result[i] = isValid(ships)
	}
	for _, v := range result {
		printResult(v)
	}
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readShips() []int {
	raw := strings.Fields(input())
	result := make([]int, len(raw))
	for i, v := range raw {
		result[i], _ = strconv.Atoi(v)
	}
	return result
}

func isValid(ships []int) bool {
	shipsCount := map[int]int{4: 1, 3: 2, 2: 3, 1: 4}
	for _, v := range ships {
		shipsCount[v]--
		if shipsCount[v] < 0 {
			return false
		}
	}
	return true
}

func printResult(isValid bool) {
	if isValid {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
