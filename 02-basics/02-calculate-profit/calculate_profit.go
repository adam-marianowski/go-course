package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate = getUserInput()
	ebt, profit, ratio := calcResults(revenue, expenses, taxRate)
	printResults(ebt, profit, ratio)
}

func getUserInput() (float64, float64, float64) {
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	return revenue, expenses, taxRate
}

func calcResults(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func printResults(ebt, profit, ratio float64) {
	fmt.Printf(`
		EBT:							%.1f
		Profit after tax:				%.1f
		ratio:							%.1f
	`, ebt, profit, ratio)
}
