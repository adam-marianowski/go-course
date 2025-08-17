package main

import "fmt"

func VariadicFunctions() {
	nums := []int{1, 10, 15}
	sum := sumup(1, 10, 15)
	anotherSum := sumup(nums...)
	fmt.Println(sum, anotherSum)
}

// ... collects all params into single slice
func sumup(nums ...int) int {
	sum := 0

	for _, val := range nums {
		sum += val
	}

	return sum
}
