package main

import "fmt"

func main() {
	// 1. Number Check
	number := 7
	number = -2
	number = 0

	// Check if the number is positive, negative, or zero
	if number > 0 {
		fmt.Println(number, "is positive.")
	} else if number < 0 {
		fmt.Println(number, "is negative.")
	} else {
		fmt.Println(number, "is zero.")
	}

	fmt.Println() // empty line for readability

	// 2. Day Planner
	dayOfWeek := "Tuesday"
	dayOfWeek = "Monday"
	dayOfWeek = "Wednesday"
	dayOfWeek = "Thursday"
	dayOfWeek = "Friday"
	dayOfWeek = "Saturday"
	dayOfWeek = "Sunday"
	dayOfWeek = "Goday"

	switch dayOfWeek {
	case "Monday":
		fmt.Println("Time to work!")
	case "Tuesday":
		fmt.Println("Taco Tuesday!")
	case "Wednesday":
		fmt.Println("Hump day!")
	case "Thursday":
		fmt.Println("Almost there!")
	case "Friday":
		fmt.Println("Weekend is coming!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend fun!")
	default:
		fmt.Println("Did the calendar get more days?!")
	}

	fmt.Println() // empty line for readability

	// 3. Countdown
	for i := 5; i >= 1; i-- {
		fmt.Println("Count is:", i)
	}
	fmt.Println("Blast off!")

	fmt.Println() // empty line for readability

	// 4. Even Numbers
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even.")
		}
	}

	fmt.Println() // empty line for readability

	// Using continue to skip printing odd numbers
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i, "is even.")
	}
}
