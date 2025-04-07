package main

import "fmt";

func main() {
	// 1. Your Info:

	// Create variables for my name and age
	myName := "Ahmad Shahzad";
	myAge := 18;
	likeGo := true;

	// Print a message using the variables with fmt.Println
	fmt.Println("Using fmt.Println:");
	fmt.Println("My name is", myName, "and I am", myAge, "years old, and it is", likeGo, "that I like Go!");
	// Print a message using the variables with fmt.Printf
	fmt.Println(""); // Print an empty line for better readability
	fmt.Printf("Using fmt.Printf:\n");
	fmt.Printf("My name is %s and I am %d years old, and it is %t that I like Go!\n", myName, myAge, likeGo);

	// 2. Simple Calculation:
	fmt.Println(""); // Print an empty line for better readability
	fmt.Println("Integer Calculation:");

	// Create variables for two numbers
	var num1 int = 15;
	var num2 int = 4;

	// Calculate the sum and product of the two numbers
	total := num1 + num2;
	product := num1 * num2;

	// Print the results using fmt.Println
	fmt.Println("num1:", num1);
	fmt.Println("num2:", num2);
	fmt.Println("Sum:", total);
	fmt.Println("Product:", product);

	// 3. Float Calculation:
	fmt.Println(""); // Print an empty line for better readability
	fmt.Println("Float Calculation:");

	// Create variables for two float numbers
	var num3 float64 = 10.5;
	var num4 float64 = 4.25;

	// Calculate the sum of the two float numbers
	totalPrice := num3 + num4;

	// Print the result using fmt.Printf
	fmt.Printf("num3: $%.2f USD\n", num3);
	fmt.Printf("num4: $%.2f USD\n", num4);
	fmt.Printf("Total Price: $%.2f USD\n", totalPrice);
}
