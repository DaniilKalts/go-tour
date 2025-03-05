/*
Loop - is a construction that repeatedly executes a block of code until a specific condition is met.
*/

package basics

import (
	"fmt"
	"strconv"
	"time"
)

func ExampleLoops() {
	// The basic for loop has three components separated by semicolons:
	// 1) the init statement: executed before the first iteration
	// 2) the condition expression: evaluated before every iteration
	// 3) the post statement: executed at the end of every iteration
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("BOOOOM")
	time.Sleep(time.Second)

	fmt.Println()
	fmt.Println("Sorry, I couldn't help but to do that !)")
	fmt.Println()

	// While-Like Loop
	age := 0
	for age < 18 {
		fmt.Println(strconv.Itoa(age) + " years old. " + "You are too young!")
		time.Sleep(time.Millisecond * 500)
		age++
	}
	fmt.Println(strconv.Itoa(age) + " years old. " + "You are young!")

	fmt.Println()
	fmt.Println("Woah, kids grow up so fast!")
	fmt.Println()

	// Infinite Loop
	fmt.Println("PULL UPS TRAINING:")
	set := 1
	for {
		if set == 2 {
			set++
			continue
		}
		if set == 4 {
			break
		}

		time.Sleep(time.Second * time.Duration(set+1))
		fmt.Printf("%d set: %d rep(s)\n", set, set)
		set++
	}
	time.Sleep(time.Second)
	fmt.Println("Did you really think I'd do Pull-Ups forever? I am not that crazy.")
	fmt.Println("And yeah, I've skipped the 2nd set.")
	fmt.Println()

	// Looping Over Collections

	// Slice or Array
	nums := []int{1, 2, 3, 4, 5}
	for idx, num := range nums {
		fmt.Printf("Index: %d, Number: %d\n", idx, num)
	}
	fmt.Println()

	// Map
	names := map[string]string{
		"a": "Abram",
		"d": "David",
		"g": "Gavriel",
		"j": "Joseph",
	}
	for start, name := range names {
		fmt.Printf("I know the name that starts with %s! That is %s.\n", start, name)
	}
}
