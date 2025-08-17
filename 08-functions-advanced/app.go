package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers2 := []int{5, 3, 5}
	doubled := transformNums(numbers, double)
	tripled := transformNums(numbers, triple)
	fmt.Println(numbers, doubled, tripled)

	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&numbers2)

	transformedNums1 := transformNums(numbers, transformFn1)
	transformedNums2 := transformNums(numbers, transformFn2)

	println(transformedNums1, transformedNums2)

}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func transformNums(numbers []int, transformFn transformFn) []int {
	dNumbers := []int{}

	for _, val := range numbers {
		dNumbers = append(dNumbers, transformFn(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
