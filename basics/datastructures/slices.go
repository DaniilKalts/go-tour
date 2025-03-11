/*
Slice - a dynamically-size, contiguous collections of elements of a single type.
*/
package datastructures

import "fmt"

func printSlice[T any](slice []T, message string) {
	fmt.Println(message)
	fmt.Printf("- Elements: %v\n", slice)
	fmt.Printf("- Length: %d\n", len(slice))
	fmt.Printf("- Capacity: %d\n", cap(slice))
}

func ExampleSlices() {
	// Declaring a nil slice of strings
	var programmingLanguages []string
	printSlice(programmingLanguages, "The slice of programming languages:")

	fmt.Println()
	fmt.Println("Wait.. I must know at least one programming language!")
	fmt.Println()

	// Appending elements to the slice
	programmingLanguages = append(programmingLanguages, "Python", "JavaScript", "TypeScript")
	programmingLanguages = append(programmingLanguages, "Go")
	printSlice(programmingLanguages, "Updated slice of programming languages:")

	fmt.Println()

	// Creating a slice from a slice
	frontendProgrammingLanguages := programmingLanguages[1:3]
	printSlice(frontendProgrammingLanguages, "The slice of frontend programming languages I've worked with:")

	fmt.Println()

	// Changing a value of a reference in a slice
	frontendProgrammingLanguages[0] = "JAVASCRIPT"
	fmt.Printf("The slice of programming languages: %v\n", programmingLanguages)
	fmt.Printf("The slice of frontend programming languages: %v\n", frontendProgrammingLanguages)

	fmt.Println()

	linuxDistributions := []string{"Ubuntu", "Fedora", "Arch"}
	printSlice(linuxDistributions, "The slice of Linux Distributions (i use arch btw):")

	fmt.Println()

	// Creating slices with make
	gradesWithMake := make([]uint8, 5)
	printSlice(gradesWithMake, "The slice of my grades:")

	fmt.Println()

	// Creating a slice of slices
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	printSlice(matrix, "The slice of matrix:")

	fmt.Println()

	// Iterating over a slice
	var pow = []uint8{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
}
