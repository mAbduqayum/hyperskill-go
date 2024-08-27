//package coffee_machine

package main

import (
	"fmt"
)

const (
	Espresso = iota + 1
	Latte
	Cappuccino
	Qandchoy
)

type BeverageRecipe struct {
	water  int
	milk   int
	teaBag int
	coffee int
	price  int
}

type CoffeeMachine struct {
	water   int
	milk    int
	coffee  int
	teaBag  int
	cups    int
	money   int
	recipes map[int]BeverageRecipe
}

func NewCoffeeMachine() *CoffeeMachine {
	return &CoffeeMachine{
		water:  400,
		milk:   540,
		teaBag: 700,
		coffee: 120,
		cups:   9,
		money:  550,
		recipes: map[int]BeverageRecipe{
			Espresso:   {water: 250, milk: 0, teaBag: 0, coffee: 16, price: 4},
			Latte:      {water: 350, milk: 75, teaBag: 0, coffee: 20, price: 7},
			Cappuccino: {water: 200, milk: 100, teaBag: 0, coffee: 12, price: 6},
			Qandchoy:   {water: 400, milk: 0, teaBag: 1, coffee: 0, price: 2},
		},
	}
}

func (m *CoffeeMachine) display() {
	fmt.Printf(`The coffee machine has:
%d ml of water
%d ml of milk
%d teaBags
%d g of coffee beans
%d disposable cups
$%d of money
`, m.water, m.milk, m.teaBag, m.coffee, m.cups, m.money)
}

func (m *CoffeeMachine) purchase() {
	fmt.Println("What do you want to buy?\n" +
		"1 - espresso\n" +
		"2 - latte\n" +
		"3 - cappuccino\n" +
		"4 - qandchoy")
	item := readValue("Enter your choice:")
	m.makeBeverage(item)
}

func (m *CoffeeMachine) makeBeverage(item int) {
	recipe, ok := m.recipes[item]
	if !ok {
		fmt.Println("Unknown beverage type!")
		return
	}

	if !m.hasEnoughResources(recipe) {
		return
	}

	m.water -= recipe.water
	m.milk -= recipe.milk
	m.teaBag -= recipe.teaBag
	m.coffee -= recipe.coffee
	m.money += recipe.price
	m.cups--
	fmt.Println("I have enough resources, making you a beverage!")
}

func (m *CoffeeMachine) hasEnoughResources(recipe BeverageRecipe) bool {
	switch {
	case m.water < recipe.water:
		fmt.Println("Sorry, not enough water!")
		return false
	case m.milk < recipe.milk:
		fmt.Println("Sorry, not enough milk!")
		return false
	case m.teaBag < recipe.teaBag:
		fmt.Println("Sorry, not enough teaBag!")
		return false
	case m.coffee < recipe.coffee:
		fmt.Println("Sorry, not enough coffee!")
		return false
	case m.cups == 0:
		fmt.Println("Sorry, not enough cups!")
		return false
	default:
		return true
	}
}

func (m *CoffeeMachine) fill() {
	m.water += readValue("Write how many ml of water you want to add:")
	m.milk += readValue("Write how many ml of milk you want to add:")
	m.teaBag += readValue("Write how many teaBags you want to add:")
	m.coffee += readValue("Write how many grams of coffee beans you want to add:")
	m.cups += readValue("Write how many disposable cups you want to add:")
}

func (m *CoffeeMachine) take() {
	fmt.Printf("I gave you $%d\n", m.money)
	m.money = 0
}

func (m *CoffeeMachine) run() {
	for {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		var action string
		fmt.Scan(&action)
		switch action {
		case "buy":
			m.purchase()
		case "fill":
			m.fill()
		case "take":
			m.take()
		case "remaining":
			m.display()
		case "exit":
			return
		}
		fmt.Println()
	}
}

func readValue(prompt string) int {
	fmt.Println(prompt)
	var value int
	fmt.Scan(&value)
	return value
}

func main() {
	machine := NewCoffeeMachine()
	machine.run()
}
