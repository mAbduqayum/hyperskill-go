package main

import (
	"fmt"
	"strings"
)

type Cinema struct {
	seats            [][]string
	purchasedTickets int
	currentIncome    int
	totalIncome      int
}

func main() {
	var rows, seats int

	fmt.Println("Enter the number of rows:")
	fmt.Scan(&rows)

	fmt.Println("Enter the number of seats in each row:")
	fmt.Scan(&seats)

	cinema := newCinema(rows, seats)

	for {
		fmt.Println("\n1. Show the seats")
		fmt.Println("2. Buy a ticket")
		fmt.Println("3. Statistics")
		fmt.Println("0. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			cinema.showSeats()
		case 2:
			cinema.buyTicket()
		case 3:
			cinema.showStatistics()
		case 0:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func newCinema(rows, seats int) *Cinema {
	c := &Cinema{
		seats: make([][]string, rows),
	}
	for i := range c.seats {
		c.seats[i] = make([]string, seats)
		for j := range c.seats[i] {
			c.seats[i][j] = "S"
		}
	}
	c.calculateTotalIncome()
	return c
}

func (c *Cinema) showSeats() {
	fmt.Println("\nCinema:")
	fmt.Print("  ")
	for i := 1; i <= len(c.seats[0]); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	for i := range c.seats {
		fmt.Printf("%d %s\n", i+1, strings.Join(c.seats[i], " "))
	}
}

func (c *Cinema) buyTicket() {
	var rowNum, seatNum int

	for {
		fmt.Println("\nEnter a row number:")
		fmt.Scan(&rowNum)

		fmt.Println("Enter a seat number in that row:")
		fmt.Scan(&seatNum)

		if rowNum < 1 || rowNum > len(c.seats) || seatNum < 1 || seatNum > len(c.seats[0]) {
			fmt.Println("Wrong input!")
			continue
		}

		if c.seats[rowNum-1][seatNum-1] == "B" {
			fmt.Println("That ticket has already been purchased!")
			continue
		}

		break
	}

	price := c.calculateTicketPrice(rowNum)
	fmt.Printf("Ticket price: $%d\n", price)

	c.seats[rowNum-1][seatNum-1] = "B"
	c.purchasedTickets++
	c.currentIncome += price
}

func (c *Cinema) calculateTicketPrice(chosenRow int) int {
	totalSeats := len(c.seats) * len(c.seats[0])
	if totalSeats <= 60 {
		return 10
	}

	frontRows := len(c.seats) / 2
	if chosenRow <= frontRows {
		return 10
	}
	return 8
}

func (c *Cinema) showStatistics() {
	totalSeats := len(c.seats) * len(c.seats[0])
	percentage := float64(c.purchasedTickets) / float64(totalSeats) * 100

	fmt.Printf("\nNumber of purchased tickets: %d\n", c.purchasedTickets)
	fmt.Printf("Percentage: %.2f%%\n", percentage)
	fmt.Printf("Current income: $%d\n", c.currentIncome)
	fmt.Printf("Total income: $%d\n", c.totalIncome)
}

func (c *Cinema) calculateTotalIncome() {
	totalSeats := len(c.seats) * len(c.seats[0])
	if totalSeats <= 60 {
		c.totalIncome = totalSeats * 10
	} else {
		frontRows := len(c.seats) / 2
		frontIncome := frontRows * len(c.seats[0]) * 10
		backIncome := (len(c.seats) - frontRows) * len(c.seats[0]) * 8
		c.totalIncome = frontIncome + backIncome
	}
}
