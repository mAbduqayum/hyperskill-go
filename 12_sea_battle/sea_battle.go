package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		ships := parseShips(scanner.Text())
		printResult(isValid(ships))
	}
}

func parseShips(input string) []int {
	raw := strings.Fields(input)
	result := make([]int, len(raw))
	for i, v := range raw {
		num, _ := strconv.Atoi(v)
		result[i] = num
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
