package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Enter the investment amount: ($) ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected annual return rate (%): ")
	fmt.Scan(&expectedReturnRate)

	fV, frV := calculateInvestment(investmentAmount, expectedReturnRate, years)

	fmt.Printf(`
	Future value:		%.1f
	With inflation:		%.1f
	`, fV, frV)

}

func calculateInvestment(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/1000, years)
	return fv, rfv
}

// EXTRA NOTES

// it does not have to be the same type.
// it can also be used with :=
// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

// another way to format strings:
// fmt.Printf(`Future Value is: %.1f
// with adjusted inflation the result is: %.1f`,
// 	futureValue, futureRealValue)

// fmt.Println("Future Value: $", futureValue) => another way
// fmt.Printf("Future value: %v\n", futureValue) => another way

// outputing values
/*
	fmt.Printf("Future Value: %.1f\n", futureValue)
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)

	fmt.Printf("Future Real Value: %.1f\n", futureRealValue)
	formattedRFV := fmt.Sprintf("Future Value: %.1f\n", futureRealValue)

	fmt.Println("=============== with sprint function ==================")
	fmt.Print(formattedFV, formattedRFV)
*/

/*another return syntax
func calculateInvestment(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/1000, years)
	return
}
*/
