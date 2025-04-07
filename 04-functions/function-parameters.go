package main

import "fmt"

// Define 'greetPerson' function that takes one parameter:
// 'name' of type string
func greetPerson(name string) {
	fmt.Println("Hello,", name+"!")
	fmt.Println("Nice to meet you.")
}

// Define 'addNumbers' function that takes two parameters:
// 'num1' of type int, 'num2' of type int
func addNumbers(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
}

// Go shortcut: If multiple parameters have the same type, you can
// write the type at the end only once
// Define 'addNumbersShort' function that takes two parameters:
// 'num1' and 'num2' of type int both
func addNumbersShort(num1, num2 int) {
	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
}

func main() {
	fmt.Println("Greet Person Function:\n")
	// Call greetPerson, passing Alice as the argument for the 'name' parameter
	greetPerson("Alice")
	fmt.Println("----------")
	greetPerson("Bob") // Call greetPerson again, but this time with Bob instead of Alice

	fmt.Println("\nAdd Numbers Function:\n")
	// Call addNumbers with 5 and 7 as arguments
	addNumbers(5, 7)
	addNumbers(10, 20) // Call addNumbers again, but for 10 and 20

	fmt.Println("\nAdd Numbers Shortcut Function:\n")
	// Call addNumbersShort with 5 adn 7 as arguments
	addNumbersShort(5, 7)
	addNumbersShort(10, 20) // Call addNumbersShort again, but for 10 and 20
	fmt.Println("\naddNumbers and addNumbersShort work the same in functionality but use a different way of adding type to their parameters.")
}
