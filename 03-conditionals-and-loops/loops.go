package main

import "fmt"

func main() {
	// Add a terminating line of sorts at the start
	fmt.Println("-------------------------")
	// Condition-Only Loop
	fmt.Println("Condition-Only Loop: While Loop Alternative")

	// Create a count variable
	count := 1

	// Create a loop that runs until count is a certain value
	for count <= 5 {
		fmt.Println("Count is:", count)
		count = count + 1
		// count++ is also valid and adds 1 to the count
	}

	// Add a terminating line of sorts after the loop ends
	fmt.Println("-------------------------")

	// Classic For Loop
	fmt.Println("Classic For Loop")

	// Create a loop that runs until 5 times
	for i := 1; i <= 5; i++ {
		fmt.Println("i is:", i)
	}

	// Add a terminating line of sorts after the loop ends
	fmt.Println("-------------------------")

	// Loop with break and continue
	fmt.Println("Loop With break And continue Keywords")

	// Create a loop with break and continue keywords
	for j := 1; j <= 10; j++ {
		// Use continue to skip printing 3 by jumping over the rest of the code
		// and going to the next iteration of the loop
		if j == 3 {
			fmt.Println("Skipping 3!")
			continue
		}

		// Use break to end the loop and not print anything further when j is equal to 8
		if j == 8 {
			fmt.Println("Stopping at 8!")
			break
		}

		// Print regular value of j
		fmt.Println("j is:", j)
	}

	// Add a terminating line of sorts after the loop ends
	fmt.Println("-------------------------")
}
