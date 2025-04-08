package main

import "fmt"

// A division function that returns multiple returns
// one for the result and one for the success/failure of the
// operation based on the value of the denominator.
func divide(numerator, denominator float64) (float64, bool) {
	// Return 0.0 for the value and false for the success status
	// if the denominator is 0.0 and the code is trying to divide by 0.0
	if denominator == 0.0 {
		return 0.0, false
	}

	// Calculate the quotient of the operation normally
	quotient := numerator / denominator

	// Return the quotient and success as true
	return quotient, true
}

func main() {
	// Call divide for values that should succeed
	q1, success1 := divide(10.0, 2.0)

	if success1 {
		fmt.Printf("10.0 / 2.0 = %.2f\n", q1)
	} else {
		fmt.Println("Error! Division by zero attempted!")
	}

	// Call divide for values that should fail
	q2, success2 := divide(4.0, 0.0)

	if success2 {
		fmt.Printf("5.0 / 0.0 = %.2f\n", q2)
	} else {
		fmt.Println("Error! Division by zero attempted!")
	}

	// Call divide but ignore the second return value for the success
	q3, _ := divide(9.0, 3.0)
	fmt.Printf("Just the quotient of 9.0 / 3.0 = %.2f\n", q3)
}
