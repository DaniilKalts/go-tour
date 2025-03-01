/*
Function - is a reusable block of code that performs a specific task.
*/

package basics

import "fmt"

func ExampleFunctions() {
	fmt.Printf("The sum of %d and %d is %d", 10, 20, Add(10, 20))
	fmt.Println()
	fmt.Println()

	fmt.Printf("The factorial of %d is %d", 5, Factorial(5))
	fmt.Println()
	fmt.Println()

	result, err := Divide(5.0, 0.0)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("The division of %f by %f is %f", 5.0, 0.0, result)
	}
	fmt.Println()
}

func Add(a, b int) int {
	return a + b
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Divide(dividend, divisor float64) (float64, error) {
	fmt.Printf("Try divide %f by %f...\n", dividend, divisor)
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by zero!")
	}
	return dividend / divisor, nil
}
