/*
Struct - a fixed collection of named fields with defined types.
*/

package datastructures

import (
	"errors"
	"fmt"
	"log"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func newUser(firstName, lastName, email, password string) (*User, error) {
	user := &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	if err := user.updatePassword(password); err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) updatePassword(newPassword string) error {
	if len(newPassword) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}
	user.Password = newPassword

	return nil
}

func ExampleStructs() {
	// Create a complete User struct using a struct literal
	person := User{FirstName: "Daniil", LastName: "Kalts", Email: "fake@gmail.com", Password: "13578642"}

	fmt.Println("Information about the person:")
	fmt.Printf("- %s\n", person.FirstName)
	fmt.Printf("- %s\n", person.LastName)
	fmt.Printf("- %s\n", person.Email)
	fmt.Printf("- %s\n", person.Password)

	fmt.Println()

	// Update the user's password via a pointer
	var personPtr *User = &person
	personPtr.Password = "12348765"

	fmt.Println("UPD: Information about the person:")
	fmt.Printf("- %s\n", person.FirstName)
	fmt.Printf("- %s\n", person.LastName)
	fmt.Printf("- %s\n", person.Email)
	fmt.Printf("- %s\n", person.Password)

	fmt.Println()
	fmt.Println("==============")
	fmt.Println()

	person2, err := newUser("John", "Doe", "johndoe@gmail.com", "12345678")
	if err != nil {
		log.Panic("Failed to create a new user")
	}

	fmt.Println("Information about the new person:")
	fmt.Printf("- %s\n", person2.FirstName)
	fmt.Printf("- %s\n", person2.LastName)
	fmt.Printf("- %s\n", person2.Email)
	fmt.Printf("- %s\n", person2.Password)

	fmt.Println()

	if err := person2.updatePassword("4321"); err != nil {
		fmt.Printf("Error updating password: %s\n\n", err)
	}

	fmt.Println("UPD: Information about the person:")
	fmt.Printf("- %s\n", person2.FirstName)
	fmt.Printf("- %s\n", person2.LastName)
	fmt.Printf("- %s\n", person2.Email)
	fmt.Printf("- %s\n", person2.Password)
}
