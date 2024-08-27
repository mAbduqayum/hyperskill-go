package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func main() {
	var paymentType string
	var payment, principal, interest float64
	var periods int

	flag.StringVar(&paymentType, "type", "", "type of payment: 'annuity' or 'diff'")
	flag.Float64Var(&payment, "payment", 0, "the monthly payment amount")
	flag.Float64Var(&principal, "principal", 0, "the loan principal amount")
	flag.IntVar(&periods, "periods", 0, "the number of months needed to repay the loan")
	flag.Float64Var(&interest, "interest", 0, "the loan interest")

	flag.Parse()

	if !validateInput(paymentType, payment, principal, periods, interest) {
		fmt.Println("Incorrect parameters")
		os.Exit(1)
	}

	monthlyInterest := interest / (12 * 100)

	switch paymentType {
	case "diff":
		if payment != 0 {
			fmt.Println("Incorrect parameters")
			os.Exit(1)
		}
		calculateDiff(principal, periods, monthlyInterest)
	case "annuity":
		calculateAnnuity(payment, principal, periods, monthlyInterest)
	default:
		fmt.Println("Incorrect parameters")
		os.Exit(1)
	}
}

func validateInput(paymentType string, payment, principal float64, periods int, interest float64) bool {
	if paymentType == "" || interest == 0 {
		return false
	}
	if payment < 0 || principal < 0 || periods < 0 || interest < 0 {
		return false
	}
	paramCount := 0
	if payment != 0 {
		paramCount++
	}
	if principal != 0 {
		paramCount++
	}
	if periods != 0 {
		paramCount++
	}
	return paramCount >= 2 // Changed from 3 to 2
}

func calculateDiff(principal float64, periods int, interest float64) {
	var totalPayment float64
	for m := 1; m <= periods; m++ {
		payment := (principal / float64(periods)) + (interest * (principal - ((principal * float64(m-1)) / float64(periods))))
		roundedPayment := math.Ceil(payment)
		totalPayment += roundedPayment
		fmt.Printf("Month %d: payment is %d\n", m, int(roundedPayment))
	}
	overpayment := totalPayment - principal
	fmt.Printf("\nOverpayment = %d\n", int(overpayment))
}

func calculatePayment(principal float64, periods int, interest float64) float64 {
	return principal * (interest * math.Pow(1+interest, float64(periods))) / (math.Pow(1+interest, float64(periods)) - 1)
}

func calculateAnnuity(payment, principal float64, periods int, interest float64) {
	if payment == 0 {
		payment = calculatePayment(principal, periods, interest)
		fmt.Printf("Your annuity payment = %d!\n", int(math.Ceil(payment)))
	} else if principal == 0 {
		principal = calculatePrincipal(payment, periods, interest)
		fmt.Printf("Your loan principal = %d!\n", int(principal))
	} else if periods == 0 {
		periods = calculatePeriods(payment, principal, interest)
		printLoanTerm(periods)
	}

	totalPayment := math.Ceil(payment) * float64(periods)
	overpayment := totalPayment - principal
	fmt.Printf("Overpayment = %d\n", int(overpayment))
}

func calculatePrincipal(payment float64, periods int, interest float64) float64 {
	return payment / ((interest * math.Pow(1+interest, float64(periods))) / (math.Pow(1+interest, float64(periods)) - 1))
}

func calculatePeriods(payment, principal, interest float64) int {
	return int(math.Ceil(math.Log(payment/(payment-interest*principal)) / math.Log(1+interest)))
}

func printLoanTerm(periods int) {
	years := periods / 12
	months := periods % 12

	if years == 0 {
		fmt.Printf("It will take %d months to repay this loan!\n", months)
	} else if months == 0 {
		if years == 1 {
			fmt.Println("It will take 1 year to repay this loan!")
		} else {
			fmt.Printf("It will take %d years to repay this loan!\n", years)
		}
	} else {
		if years == 1 {
			fmt.Printf("It will take 1 year and %d months to repay this loan!\n", months)
		} else {
			fmt.Printf("It will take %d years and %d months to repay this loan!\n", years, months)
		}
	}
}
