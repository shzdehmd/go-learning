package main

import "fmt";

func main() {
	// Declaring a variable named "message" that can hold characters and words: string
	// Assigning a value to the message variable using Go short assignment
	message := "Learning Go variables!";

	// Printing the value of the message variable
	fmt.Println("Message:", message);

	// Changing the value of the message variable
	message = "Learning Go variables is fun!";

	// Printing the new value of the message variable
	fmt.Println("Updated message:", message);
}
