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

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solveTestCase()
	}
}

func solveTestCase() {
	var n int
	fmt.Fscan(in, &n)

	if solve(n) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

func solve(n int) bool {
	namePrice := make(map[string]int, n)
	priceName := make(map[int]string, n)
	for range n {
		var name string
		var price int
		fmt.Fscan(in, &name, &price)
		namePrice[name] = price
		priceName[price] = name
	}

	_, _ = in.ReadString('\n')
	var userResult string
	fmt.Fscanln(in, &userResult)
	itemsString := strings.Split(userResult, ",")
	if len(itemsString) != len(priceName) {
		return false
	}
	seenPrice := make(map[int]bool, n)
	for _, item := range itemsString {
		itemParts := strings.Split(item, ":")
		if len(itemParts) != 2 {
			return false
		}

		name := itemParts[0]
		if !isValidName(name) {
			return false
		}
		if _, exists := namePrice[name]; !exists {
			return false
		}

		priceStr := itemParts[1]
		if len(priceStr) > 1 && priceStr[0] == '0' {
			return false
		}
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			return false
		}
		if price < 1 || price > 1e9 {
			return false
		}
		if seenPrice[price] {
			return false
		}
		if price != namePrice[name] {
			return false
		}
		seenPrice[price] = true
	}
	return true
}

func isValidName(name string) bool {
	return len(name) >= 1 && len(name) <= 10
}
