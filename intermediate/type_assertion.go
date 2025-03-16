/*
Type Assertion - extracts the actual value from an interface
*/

package intermediate

import "fmt"

type ProgrammingLanguage struct {
	Name      string
	Creator   string
	Paradigms []string
}

func ExampleTypeAssertions() {
	stuff := []interface{}{"bottle", 99.99, 999}

	for _, item := range stuff {
		switch v := item.(type) {
		case string:
			fmt.Printf("String: %s\n", v)
		case float64:
			fmt.Printf("Float: %f\n", v)
		case int:
			fmt.Printf("Integer: %d\n", v)
		}
	}
	fmt.Println()

	python := ProgrammingLanguage{
		Name:      "Python",
		Creator:   "Guido van Rossum",
		Paradigms: []string{"imperative", "functional", "procedural", "object-oriented"},
	}

	var lang interface{} = python
	if programmingLanguage, ok := lang.(ProgrammingLanguage); ok {
		fmt.Printf("Successfully asserted to ProgrammingLanguage: %+v\n", programmingLanguage)
	} else {
		fmt.Println("Assertion to ProgrammingLanguage failed")
	}

	var langPtr interface{} = &python
	if programmingLanguagePtr, ok := langPtr.(*ProgrammingLanguage); ok {
		fmt.Printf("Successfully asserted to *ProgrammingLanguage: %+v\n", programmingLanguagePtr)
	} else {
		fmt.Println("Assertion to *ProgrammingLanguage failed")
	}
}
