package main

import "fmt"

func RecursionLesson() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(n int) int {
	// result := 1
	// for i := 1; i <= n; i++ {
	// 	result = result * i
	// }
	// return result
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

// factorial of 5 => 5*4*3*2*1=120
