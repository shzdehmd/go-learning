package main

import "fmt"

func main() {
	// Print a message based on the day
	day := "Monday"
	day = "Friday"
	day = "Saturday"
	day = "Wednesday"

	fmt.Printf("It's %s today!\n", day)

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!!! :)")
	case "Monday":
		fmt.Println("The work week starts today. :(")
	case "Friday":
		fmt.Println("It's almost the weekend! >.<")
	default:
		fmt.Println("Just a regular weekday. O-O")
	}

	// Evaluate expressions in switch
	score := 75
	score = 91
	score = 85
	score = 60
	score = 50

	fmt.Printf("\nScore: %d\n", score)

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
}
