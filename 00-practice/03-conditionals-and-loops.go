package main

import "fmt"

func main() {
	// 1. Number Check
	number := 7
	number = -2
	number = 0

	// Check if the number is positive, negative, or zero
	if number > 0 {
		fmt.Println(number, "is positive.")
	} else if number < 0 {
		fmt.Println(number, "is negative.")
	} else {
		fmt.Println(number, "is zero.")
	}
}
