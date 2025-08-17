package main

import "fmt"

func call() {

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings {
		fmt.Println(key)
	}
}
