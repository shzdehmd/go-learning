package main

import "fmt"

// Define 'calculateSum' which takes two int parameters
// and returns one value of type int
func calculateSum(num1, num2 int) int {
	sum := num1 + num2
	return sum // Send the calculated sum back
	// Any code after return will not be executed inside this function
}

// Define 'createGreeting' which takes a name (string) parameter
// and returns a greeting (string) value
func createGreeting(name string) string {
	greeting := "Hello, " + name + "! Welcome back."
	return greeting
}

func main() {
	fmt.Println("Calculate Sum Function:\n")
	// Call calcuateSum and store the returned value in a variable
	result1 := calculateSum(10, 20)
	fmt.Println("The first sum is:", result1)

	result2 := calculateSum(50, 70)
	fmt.Println("The second sum is:", result2)

	// Directly use the function inside fmt.Println
	fmt.Println("Direct sum:", calculateSum(100, 200))

	fmt.Println("----------\n")

	fmt.Println("Create Greeting Function:\n")
	// Call createGreeting and store the returned value in a variable
	greeting1 := createGreeting("Alice")
	fmt.Println(greeting1)

	greeting2 := createGreeting("Bob")
	fmt.Println(greeting2)

	// Directly use the function inside fmt.Println
	fmt.Println(createGreeting("Jack"))

	fmt.Println("----------\n")
}
