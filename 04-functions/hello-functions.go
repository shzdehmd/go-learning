package main

import "fmt"

func sayHello() {
	fmt.Println("Hello there!")
	fmt.Println("Welcome to the world of functions")
}

func main() {
	fmt.Println("Starting the program...")
	fmt.Println() // empty line for readability

	// Call the sayHello function
	sayHello()

	fmt.Println("Called sayHello function once.")
	fmt.Println() // empty line for readability

	// Call the sayHello function again
	sayHello()

	fmt.Println("Called the sayHello function once more.")
	fmt.Println() // empty line for readability

	fmt.Println("Program finished.")
}
