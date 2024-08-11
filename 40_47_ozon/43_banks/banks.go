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
		processTest()
	}
}

func processTest() {
	var maxDollars float64
	var exchangeRates [3][6]int

	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &exchangeRates[i][0], &exchangeRates[i][1])
		fmt.Fscan(in, &exchangeRates[i][2], &exchangeRates[i][3])
		fmt.Fscan(in, &exchangeRates[i][1], &exchangeRates[i][2])
		fmt.Fscan(in, &exchangeRates[i][4], &exchangeRates[i][5])
		fmt.Fscan(in, &exchangeRates[i][3], &exchangeRates[i][2])
		fmt.Fscan(in, &exchangeRates[i][5], &exchangeRates[i][4])
	}

	// Get the best ruble to dollar exchange rate
	bestRubleToDollar := 0.0
	for i := 0; i < 3; i++ {
		rubleToDollar := 1.0 / float64(exchangeRates[i][0])
		if rubleToDollar > bestRubleToDollar {
			bestRubleToDollar = rubleToDollar
		}
	}

	// Get the best ruble to euro to dollar exchange rate
	bestRubleEuroDollar := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j {
				rubleToEuro := 1.0 / float64(exchangeRates[i][2])
				euroToDollar := 1.0 / float64(exchangeRates[j][5])
				rubleEuroDollar := rubleToEuro * euroToDollar
				if rubleEuroDollar > bestRubleEuroDollar {
					bestRubleEuroDollar = rubleEuroDollar
				}
			}
		}
	}

	fmt.Fprintln(out, bestRubleEuroDollar, bestRubleToDollar)
	// Get the maximum dollars obtained
	maxDollars = bestRubleToDollar
	if bestRubleEuroDollar > maxDollars {
		maxDollars = bestRubleEuroDollar
	}

	fmt.Fprintf(out, "%.6f\n", maxDollars)
}
