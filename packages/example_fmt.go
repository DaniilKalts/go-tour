package packages

import "fmt"

func ExampleFmt() {
	// fmt.Print - outputs without extra spacing or new lines.
	fmt.Print("Hello, ")
	fmt.Print("World!\n")
	fmt.Print("This is ", "the most iconic ", "phrase in the programming world, right?)\n")

	// fmt.Println - outputs with spaces between arguments and a new line.
	fmt.Println()
	fmt.Println("Hello,", "World!")
	fmt.Println("Same result with less efforts!")
	fmt.Println()

	// fmt.Printf - formats output without a new line.
	// %s - string
	// %d - decimal integer
	fmt.Printf("My name is %s and I am %d years old.\n", "Daniil", 18)

	// fmt.Sprintf - returns a formatted string.
	greeting := fmt.Sprintf("My name is %s and I am %d years old.", "Daniil", 18)
	fmt.Println("Greeting:", greeting)
}
