package main

import "fmt"

// Calculator function
func calculate(operator string, x, y float64) float64 {
	var result float64

	// Print the current operation in order.
	fmt.Printf("Operation: %.2f %s %.2f\n", x, operator, y)

	switch operator {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		// Show an error when trying to divide by 0.
		if y == 0.0 {
			fmt.Println("Cannot divide by zero!")
			result = 0.0
		} else {
			result = x / y
		}
	case "%":
		// Show an error when trying to calculate modulo with 0 denominator.
		if y == 0.0 {
			fmt.Println("Cannot divide by zero!")
			result = 0.0
		} else {
			// Convert the inputs to int for modulus and then convert the result to float64
			result = float64(int(x) % int(y))
		}
	default:
		// If any other operator is passed, return an error.
		fmt.Println("Invalid operation!")
		result = 0.0
	}

	return result
}

// FizzBuzz function
func fizzbuzz(limit int) {
	if limit < 0 { // Return an error if the user passes 0 or a negative number
		fmt.Println("Uh, oh! You cannot do that! Pass a limit greater than 0.")
	} else {
		fmt.Println("FizzBuzz running until number:", limit) // Show the limit

		for i := 0; i <= limit; i++ {
			if i%3 == 0 && i%5 == 0 { // FizzBuzz if divisible by 3 and 5 perfectly
				fmt.Println(i, "is FizzBuzz!")
			} else if i%3 == 0 { // Fizz if only divisible by 3
				fmt.Println(i, "is Fizz!")
			} else if i%5 == 0 { // Buzz if only divisible by 5
				fmt.Println(i, "is Buzz!")
			} else { // Dud if not divisible by 3 and 5
				fmt.Println(i, "is a dud!")
			}
		}
	}
}

// Function to create a new string with repetition of the original
func stringRepeater(text string, count int) string {
	// Give an error if the count is less than or equal to zero.
	if count < 1 {
		return "" // Return empty string if the repeat count is 0 or less
	} else {
		result := ""
		for i := 0; i < count; i++ { // Repeat the loop for as many times as the count says
			result += text + "\n" // Append the string to the result with a newline at the end for better readability
		}
		return result // Return the final result
	}
}

// Function to generate formal and informal greetings
func generateGreetings(name string) (string, string) {
	if name == "" { // Check if the name is empty
		return "", "" // Return empty greetings
	} else {
		formalGreeting := "Dear " + name + ","  // Create a formal greeting as Dear name,
		informalGreeting := "Hey " + name + "!" // Create an informal greeting as Hey name!

		return formalGreeting, informalGreeting // Return the greetings
	}
}

func main() {
	// Testing the Calculator function
	fmt.Println("Result:", calculate("+", 20, 30), "\n") // Addition
	fmt.Println("Result:", calculate("-", 20, 30), "\n") // Subtraction
	fmt.Println("Result:", calculate("*", 20, 30), "\n") // Multiplication
	fmt.Println("Result:", calculate("/", 20, 30), "\n") // Division
	fmt.Println("Result:", calculate("/", 20, 0), "\n")  // Division by 0, should throw error and return 0
	fmt.Println("Result:", calculate("%", 50, 20), "\n") // Modulus
	fmt.Println("Result:", calculate("%", 50, 0), "\n")  // Modulus by 0, should throw error and return 0
	fmt.Println("Result:", calculate("&", 20, 30), "\n") // Invalid operator, should throw error and return 0

	// Testing the FizzBuzz function
	fizzbuzz(-10) // Test for a negative number

	fmt.Println()
	fizzbuzz(20) // Test for a limit up to 20
	fmt.Println()

	// Testing the stringRepeater function
	fmt.Println(stringRepeater("Hello, world!", 10)) // Repeat 10 times
	fmt.Println()
	fmt.Println(stringRepeater("Hello, world!", 0)) // Repeat 0 times, should be empty

	// Testing the generateGreetings function
	formal1, informal1 := generateGreetings("") // Test with empty string
	fmt.Printf("Formal Greeting: %s\nInformal Greeting: %s\n\n", formal1, informal1)

	formal2, informal2 := generateGreetings("Denise") // Test with a name
	fmt.Printf("Formal Greeting: %s\nInformal Greeting: %s\n\n", formal2, informal2)

	formal3, _ := generateGreetings("Denise") // Ignore the second return value
	fmt.Printf("Formal Greeting: %s\n\n", formal3)

	_, informal3 := generateGreetings("Denise") // Ignore the first return value
	fmt.Printf("Informal Greeting: %s\n", informal3)
}
