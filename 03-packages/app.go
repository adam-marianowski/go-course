package main

import (
	"bank/app/fileops"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

var balance, balanceError = fileops.GetBalanceFromFile()

func main() {
	fmt.Println("Welcome to Go-Bank!")
	fmt.Println(randomdata.PhoneNumber())

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
			fileops.WriteBalanceToFile(balance)
		case 2:
			amount := getAmount()
			deposit(amount)
			fileops.WriteBalanceToFile(balance)
		case 3:
			amount := getAmount()
			withdraw(amount)
			fileops.WriteBalanceToFile(balance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using Go-Bank!")
			return
		}
	}
}
