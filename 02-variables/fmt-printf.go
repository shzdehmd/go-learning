package main

import "fmt";

func main() {
	// Create a variable to store name and age and height
	name := "Jake Daniels";
	age := 65;
	height := 1.65;

	// Print the name, age and height using fmt.Printf
	// Printf needs a new line terminator at the end as \n
	// %s is for strings
	// %d is for integers
	// %f is for float
	fmt.Printf("Name: %s\n", name);
	fmt.Printf("Age: %d years old\n", age);
	fmt.Printf("Height: %f meters\n", height);

	// Print the name, age and height using fmt.Printf with a message
	fmt.Printf("Hello, my name is %s, I am %d years old, and my height is %f meters.\n", name, age, height);
}
