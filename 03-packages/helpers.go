package main

import (
	"fmt"
)

func getChoice() int {
	fmt.Printf(`
_______________________
Select operation:
1. check balance
2. Deposit
3. Withdraw
4. Exit
_______________________
	`)

	var choice int
	fmt.Print("Option: ")
	fmt.Scan(&choice)

	return choice
}

func getAmount() int {
	var amount int
	fmt.Print("Amount:")
	fmt.Scan(&amount)
	return amount
}

func withdraw(amount int) {

	if amount <= 0 {
		fmt.Println("Invalid amount. It must be greater than 0")
		return
	}

	if balance-float64(amount) < 0 {
		fmt.Println("Not enough balance")
		return
	}

	balance -= float64(amount)
	fmt.Printf("new balance: %.2f\n", balance)
}

func deposit(amount int) {
	balance += float64(amount)
	fmt.Printf("new balance: %.2f\n", balance)
}
