/*
Pointer - is a variable that stores memory address of a value.
*/

package basics

import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
	age  uint8
}

// Check if the provided string (via pointer) is a palindrome
func isPalindrome(str *string) bool {
	s := strings.ToLower(*str)

	l := 0
	r := len(*str) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l += 1
		r -= 1
	}

	return true
}

// Modify the age of a person using a pointer receiver
func updateAge(p *Person, newAge uint8) {
	p.age = newAge
}

func ExamplePointers() {
	// & operator generates a pointer to its operand
	var age uint8 = 18
	fmt.Printf("Information about the age variable (%d):\n", age)
	fmt.Printf("- Value: %d\n", age)
	fmt.Printf("- Address: %p\n", &age)

	fmt.Println()

	// *T is a pointer to a T value
	// * operator can also retrieve a value from memory address
	var agePtr1 *uint8 = &age
	fmt.Printf("Information about the memory address (%p):\n", agePtr1)
	fmt.Printf("- Value: %d\n", *agePtr1)
	fmt.Printf("- Address: %p\n", agePtr1)

	fmt.Println()

	// Nil pointer
	var agePtr2 *uint8
	fmt.Printf("Information about the memory address (%p):\n", agePtr2)
	fmt.Println("- Value: nil pointers don't reference to a value")
	fmt.Printf("- Address: %p\n", agePtr2)

	fmt.Println()

	// Check if a name is palindrome using a pointer to a string
	var name = "Bob"
	fmt.Printf("Is the name \"%s\" palindrome? This is %t\n", name, isPalindrome(&name))

	fmt.Println()

	// new - allocate a pointer to a zero-initialized variable
	var agePtr3 = new(uint8)
	fmt.Printf("Information about the memory address (%p):\n", agePtr3)
	fmt.Printf("- Value: %d\n", *agePtr3)
	fmt.Printf("- Address: %p\n", agePtr3)

	fmt.Println()

	// Pointers with structs
	var me Person = Person{Name: "Daniil", age: 18}
	fmt.Printf("Information about the memory address (%p):\n", &me)
	fmt.Printf("- Value: %+v\n", me)
	fmt.Printf("- Address: %p\n", &me)

	updateAge(&me, 20)

	fmt.Println()
	fmt.Println("Two years later...")
	fmt.Println()

	fmt.Printf("Information about the memory address (%p):\n", &me)
	fmt.Printf("- Value: %+v\n", me)
	fmt.Printf("- Address: %p\n", &me)
}
