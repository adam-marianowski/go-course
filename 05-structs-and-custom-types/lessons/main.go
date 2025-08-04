package main

import "fmt"

type customString string

func (text customString) log() {
	fmt.Println(text)
}

func doSomething() {
	var name customString = "James"
	name.log()
}
