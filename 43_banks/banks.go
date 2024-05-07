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

type ExchangeRate struct {
	rubleToDollar float64
	rubleToEuro   float64
	dollarToRuble float64
	dollarToEuro  float64
	euroToRuble   float64
	euroToDollar  float64
}

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		processTest()
	}
}

func processTest() {
	var banks [3]ExchangeRate
	for i := 0; i < 3; i++ {
		banks[i] = ExchangeRate{
			rubleToDollar: getProportion(),
			rubleToEuro:   getProportion(),
			dollarToRuble: getProportion(),
			dollarToEuro:  getProportion(),
			euroToRuble:   getProportion(),
			euroToDollar:  getProportion(),
		}
	}

	bestRubleToDollar := best(
		banks[0].rubleToDollar,
		banks[1].rubleToDollar,
		banks[2].rubleToDollar,
	)

	bestRubleEuroDollar := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j {
				rubleToEuro := banks[i].rubleToEuro
				euroToDollar := banks[j].euroToDollar
				bestRubleEuroDollar = best(
					bestRubleEuroDollar,
					rubleToEuro*euroToDollar,
				)
			}
		}
	}

	bestRubleEuroRubleDollar := 0.0
	bestRubleDollarRubleDollar := 0.0
	bestRubleDollarEuroDollar := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if i != j && j != k && i != k {
					bestRubleEuroRubleDollar = best(
						bestRubleEuroRubleDollar,
						banks[i].rubleToEuro*banks[j].euroToRuble*banks[k].rubleToDollar,
					)
					bestRubleDollarRubleDollar = best(
						bestRubleDollarRubleDollar,
						banks[i].rubleToDollar*banks[j].dollarToRuble*banks[k].rubleToDollar,
					)
					bestRubleDollarEuroDollar = best(
						bestRubleDollarEuroDollar,
						banks[i].rubleToDollar*banks[j].dollarToEuro*banks[k].euroToDollar,
					)
				}
			}
		}
	}

	rez := best(
		bestRubleToDollar,
		bestRubleEuroDollar,
		bestRubleEuroRubleDollar,
		bestRubleDollarRubleDollar,
		bestRubleDollarEuroDollar,
	)
	fmt.Fprintf(out, "%g\n", rez)
}

func best(nums ...float64) float64 {
	rez := nums[0]
	for _, num := range nums {
		if num > rez {
			rez = num
		}
	}
	return rez
}

func getProportion() float64 {
	var a, b float64
	fmt.Fscan(in, &a, &b)
	return b / a
}
