/*
Function - is a reusable block of code that performs a specific task.
*/

package basics

import "fmt"

func ExampleFunctions() {
	fmt.Printf("The sum of %d and %d is %d", 10, 20, add(10, 20))
	fmt.Println()
	fmt.Println()

	fmt.Printf("The factorial of %d is %d", 5, factorial(5))
	fmt.Println()
	fmt.Println()

	result, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("The division of %f by %f is %f", 5.0, 0.0, result)
	}
	fmt.Println()
}

func add(a, b int) int {
	return a + b
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func divide(dividend, divisor float64) (float64, error) {
	fmt.Printf("Try divide %f by %f...\n", dividend, divisor)
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by zero!")
	}
	return dividend / divisor, nil
}
