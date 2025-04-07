package main

import "fmt"

func main() {
	// Simple age check
	age := 30
	age = 12
	age = 18

	fmt.Printf("Age: %d\n", age)

	// Check if someone is an adult or a minor based on their age
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	fmt.Println() // Empty line for better readability

	// Another example using temperature.
	temperature := 15
	temperature = 5
	temperature = 30

	fmt.Printf("Temperature: %d\n", temperature)

	if temperature >= 25 {
		fmt.Println("It's hot!")
	} else if temperature <= 10 {
		fmt.Println("It's cold!")
	} else {
		fmt.Println("The weather's nice.")
	}
}
