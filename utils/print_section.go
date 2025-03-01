package utils

import "fmt"

func PrintLibrarySection(title string, sectionFunc func()) {
	divider := "\n=================================\n"

	fmt.Println(divider)
	fmt.Printf("Example of the built-in %s library.\n\n", title)

	sectionFunc()
}

func PrintSyntaxSection(title string, sectionFunc func()) {
	divider := "\n=================================\n"

	fmt.Println(divider)
	fmt.Printf("Example of the syntax for %s.\n\n", title)

	sectionFunc()
}
