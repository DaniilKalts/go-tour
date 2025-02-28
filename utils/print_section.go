package utils

import "fmt"

func PrintSection(title string, sectionFunc func()) {
	divider := "\n=================================\n"

	fmt.Println(divider)
	fmt.Printf("Example of built-in %s library.\n\n", title)

	sectionFunc()
}
