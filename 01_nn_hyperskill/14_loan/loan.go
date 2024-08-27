package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

type LoanCalculator struct {
	PaymentType string
	Payment     float64
	Principal   float64
	Periods     int
	Interest    float64
}

func main() {
	lc := parseFlags()

	if !lc.validateInput() {
		fmt.Println("Incorrect parameters")
		os.Exit(1)
	}

	monthlyInterest := lc.Interest / (12 * 100)

	switch lc.PaymentType {
	case "diff":
		if lc.Payment != 0 {
			fmt.Println("Incorrect parameters")
			os.Exit(1)
		}
		lc.calculateDiff(monthlyInterest)
	case "annuity":
		lc.calculateAnnuity(monthlyInterest)
	default:
		fmt.Println("Incorrect parameters")
		os.Exit(1)
	}
}

func parseFlags() *LoanCalculator {
	lc := &LoanCalculator{}
	flag.StringVar(&lc.PaymentType, "type", "", "type of payment: 'annuity' or 'diff'")
	flag.Float64Var(&lc.Payment, "payment", 0, "the monthly payment amount")
	flag.Float64Var(&lc.Principal, "principal", 0, "the loan principal amount")
	flag.IntVar(&lc.Periods, "periods", 0, "the number of months needed to repay the loan")
	flag.Float64Var(&lc.Interest, "interest", 0, "the loan interest")
	flag.Parse()
	return lc
}

func (lc *LoanCalculator) validateInput() bool {
	if lc.PaymentType == "" || lc.Interest == 0 {
		return false
	}
	if lc.Payment < 0 || lc.Principal < 0 || lc.Periods < 0 || lc.Interest < 0 {
		return false
	}
	paramCount := 0
	if lc.Payment != 0 {
		paramCount++
	}
	if lc.Principal != 0 {
		paramCount++
	}
	if lc.Periods != 0 {
		paramCount++
	}
	return paramCount >= 2
}

func (lc *LoanCalculator) calculateDiff(interest float64) {
	var totalPayment float64
	for m := 1; m <= lc.Periods; m++ {
		payment := (lc.Principal / float64(lc.Periods)) + (interest * (lc.Principal - ((lc.Principal * float64(m-1)) / float64(lc.Periods))))
		roundedPayment := math.Ceil(payment)
		totalPayment += roundedPayment
		fmt.Printf("Month %d: payment is %d\n", m, int(roundedPayment))
	}
	overpayment := totalPayment - lc.Principal
	fmt.Printf("\nOverpayment = %d\n", int(overpayment))
}

func (lc *LoanCalculator) calculateAnnuity(interest float64) {
	if lc.Payment == 0 {
		lc.Payment = lc.calculatePayment(interest)
		fmt.Printf("Your annuity payment = %d!\n", int(math.Ceil(lc.Payment)))
	} else if lc.Principal == 0 {
		lc.Principal = lc.calculatePrincipal(interest)
		fmt.Printf("Your loan principal = %d!\n", int(lc.Principal))
	} else if lc.Periods == 0 {
		lc.Periods = lc.calculatePeriods(interest)
		lc.printLoanTerm()
	}

	totalPayment := math.Ceil(lc.Payment) * float64(lc.Periods)
	overpayment := totalPayment - lc.Principal
	fmt.Printf("Overpayment = %d\n", int(overpayment))
}

func (lc *LoanCalculator) calculatePayment(interest float64) float64 {
	return lc.Principal * (interest * math.Pow(1+interest, float64(lc.Periods))) / (math.Pow(1+interest, float64(lc.Periods)) - 1)
}

func (lc *LoanCalculator) calculatePrincipal(interest float64) float64 {
	return lc.Payment / ((interest * math.Pow(1+interest, float64(lc.Periods))) / (math.Pow(1+interest, float64(lc.Periods)) - 1))
}

func (lc *LoanCalculator) calculatePeriods(interest float64) int {
	return int(math.Ceil(math.Log(lc.Payment/(lc.Payment-interest*lc.Principal)) / math.Log(1+interest)))
}

func (lc *LoanCalculator) printLoanTerm() {
	years := lc.Periods / 12
	months := lc.Periods % 12

	switch {
	case years == 0:
		fmt.Printf("It will take %d months to repay this loan!\n", months)
	case months == 0:
		if years == 1 {
			fmt.Println("It will take 1 year to repay this loan!")
		} else {
			fmt.Printf("It will take %d years to repay this loan!\n", years)
		}
	default:
		if years == 1 {
			fmt.Printf("It will take 1 year and %d months to repay this loan!\n", months)
		} else {
			fmt.Printf("It will take %d years and %d months to repay this loan!\n", years, months)
		}
	}
}
