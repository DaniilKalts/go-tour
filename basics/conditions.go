/*
Condition - is a construction that evaluates a Boolean expression to determine the flow of execution in a program.
*/

package basics

import (
	"fmt"
	"strings"
)

func dayOfTheWeekMessage(day string) string {
	day = strings.TrimSpace(day)
	day = strings.ToLower(day)
	var message string

	switch day {
	case "monday":
		message = "AHH SHOOT! One more monday..."
	case "tuesday":
		message = "It's okay."
	case "wednesday":
		message = "I NEED TO BE IN CLASS BEFORE CALCULUS TEACHER!!!!!"
	case "thursday":
		message = "Phew, the hardest part of the week is over"
	case "friday":
		message = "Yay, no classes!"
	case "saturday":
		message = "Yay, today I gonna walk!"
	case "sunday":
		message = "A long stretch of darkness is coming..."
	default:
		message = "Unknown day of the week."
	}

	return message
}

func ExampleConditions() {
	var age int
	fmt.Print("What is your age? ")
	fmt.Scanln(&age)

	if age <= 4 {
		fmt.Println("You are a baby!")
	} else if age <= 12 {
		fmt.Println("You are a kid!")
	} else if age <= 17 {
		fmt.Println("You are a teen!")
	} else if age <= 120 {
		fmt.Println("You are an adult!")
	} else {
		fmt.Println("Wait, you are still alive?!")
	}

	fmt.Println()

	var day string
	fmt.Print("Enter a day of the week: ")
	fmt.Scanln(&day)

	fmt.Println(dayOfTheWeekMessage(day))
}
