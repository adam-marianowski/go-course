package main

import "fmt"

type Product struct {
	Title stringW
	ID    string
	Price float64
}

func main() {
	// ARRAYS
	prices := [4]float64{19.99, 2.99, 3.99, 1.99}
	// dynamic array
	dynamicSlices := []float64{10.99, 23.99, 9.11}

	var productNames [4]string
	productNames = [4]string{"p1", "p2", "p3", "p4"}

	fmt.Println(prices[0])

	// first value is included
	// last value is excluded
	// pricesSlice := prices[1:3]

	// we can merge  arrays with spread operator

	// MAPS
	// map = obj with kv pairs
	websites := map[string]string{
		"google": "google.com",
		"AWS":    "aws.com",
	}

	fmt.Println(websites["AWS"])

	// creating entries
	websites["Azure"] = "portal.azure.com"

	delete(websites, "AWS")

	// MAP VS STRUCT
	/*
		- maps - anything can be a key.
		- struct - predefined data structure. Can't be modified
	*/

	// MAKE
	// predefine size of the array
	// must have length defined
	// we can access undefined values like this
	// userNames[0] = "Julie"
	userNames := make([]string, 2, 5)
	userNames = append(userNames, "Adam")
	userNames = append(userNames, "Max")

	// make can be used for maps
	// we can pass only 1 arg to map
	courses := make(map[string]float64, 4)
	courses["go"] = 4.7
	courses["react"] = 4.8

	// type alias
	type courseRating map[string]float64
	courseRatings := make(courseRating, 4)
	fmt.Println(courseRatings.output())
}

// output method for courseRating
func (m courseRating) output() string {
	return fmt.Sprintf("%v", m)
}
