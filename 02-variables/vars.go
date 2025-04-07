package main

import "fmt";

func main() {
	fmt.Println("I am learning variables in Go!");

	// Declaring a variable with the var keyword
	// var VARIABLE_NAME VARIABLE_TYPE = VARIABLE_VALUE

	// Declaring a variable named "message" that can hold characters and words: string
	var message string;

	// Assigning a value separately to the message variable
	message = "Learning Go variables!";

	// Printing the value of the message variable
	fmt.Println(message);

	// Declaring a variable named "number" that can hold whole numbers: int
	// Assigning a value to the number variable in the same line it is initialized
	var number int = 10;

	// Printing the value of the number variable
	fmt.Println(number);

	// Declaring another variable named "number2" that can hold whole numbers: int
	// Not assigning a value to the number2 variable in the same line it is initialized
	var number2 int;

	// Assigning a value to the number2 variable separately
	number2 = 20;

	// Printing the value of the number2 variable
	fmt.Println(number2);
}
