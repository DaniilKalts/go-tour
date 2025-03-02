/*
Variable - is a named memory location for storing a value(s).
*/

package basics

import "fmt"

func ExampleVariables() {
	// 1) Explicit declaration: var variableName type = value
	var isAdult = true
	var isProgrammer, speaksEnglish bool = true, true

	fmt.Println("Hey, guys!")
	fmt.Println("My name is Daniil and below are a couple of facts about me:")
	fmt.Println()

	fmt.Print("Am I adult? ")
	if isAdult {
		fmt.Printf("Yes, that's %t", speaksEnglish)
	} else {
		fmt.Printf("No, that's %t", speaksEnglish)
	}
	fmt.Println()

	fmt.Print("Am I a programmer? ")
	if isProgrammer {
		fmt.Printf("Yes, that's %t", isProgrammer)
	} else {
		fmt.Printf("No, that's %t", isProgrammer)
	}
	fmt.Println()

	fmt.Print("Do I speak English? ")
	if speaksEnglish {
		fmt.Printf("Yes, that's %t", speaksEnglish)
	} else {
		fmt.Printf("No, that's %t", speaksEnglish)
	}
	fmt.Println()
	fmt.Println()

	// 2) Implicit declaration: variableName := value
	// Befare: it only works inside of a function
	name := "Daniil"
	age := 18

	fmt.Printf("As I've said my name is %s, but did you know that I am %d years old?", name, age)
	fmt.Println()
	fmt.Println()

	// 3) Constant declaration: const variableName = value
	// And yeah, you can't change a constant variable
	const pi = 3.14

	fmt.Printf("Believe or not, but PI is equal to %f", pi)
	fmt.Println()
	fmt.Println()

	// 4) Variable Declaration without initialization: var variableName type
	// They will take a Falsy value from the data type
	var count int
	fmt.Printf("Initially the counter is equal to %d", count)
	fmt.Println()
}
