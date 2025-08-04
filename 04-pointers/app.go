package main

import "fmt"

func main() {
	age := 32

	var agePointer *int = &age

	fmt.Println(*agePointer)

	getAdultYears(agePointer)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
