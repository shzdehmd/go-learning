package main

import "fmt"

func testFunction() {
	boundToFunction := "I only work inside my function"
	fmt.Println(boundToFunction)
}

func main() {
	testFunction()

	fmt.Println(boundToFunction) // this will throw an error as boundToFunction is only available in testFunction
}
