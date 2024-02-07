package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		if isValid(scanner.Text()) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func isValid(dateString string) bool {
	daysInMonth := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	fields := strings.Fields(dateString)
	d, _ := strconv.Atoi(fields[0])
	m, _ := strconv.Atoi(fields[1])
	y, _ := strconv.Atoi(fields[2])
	if m == 2 && isLeapYear(y) {
		daysInMonth[2] = 29
	}
	return d > 0 && d <= daysInMonth[m]
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func isValidBuiltin(dateString string) bool {
	_, err := time.Parse("2 1 2006", dateString)
	return err == nil
}
