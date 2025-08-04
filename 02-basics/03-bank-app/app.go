package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var balance, balanceError = getBalanceFromFile()

func main() {
	fmt.Print("Welcome to Go-Bank!")

	if balanceError != nil {
		fmt.Println("error")
		fmt.Println(balanceError)
		fmt.Println("-----------------")
		panic("Could not execute this app...")
		// or return // this will exit app.
	}

	for {
		choice := getChoice()
		switch choice {
		case 1:
			fmt.Printf("Balance: %.2f\n", balance)
			writeBalanceToFile(balance)
		case 2:
			amount := getAmount()
			deposit(amount)
			writeBalanceToFile(balance)
		case 3:
			amount := getAmount()
			withdraw(amount)
			writeBalanceToFile(balance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go-Bank!")
			return
		}
	}
}

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
	println("balance printed...")
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to read the file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse balance")
	}

	return balance, nil
}

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
