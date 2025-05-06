/*
Stringer - the interface fmt looks for when printing a value.
*/
package intermediate

import "fmt"

type Human struct {
	FirstName string
	LastName  string
	Age       int8
	Salary    int
}

type Man struct {
	Human
}

func NewMan(firstName string, lastName string, age int8, salary int) *Man {
	return &Man{
		Human: Human{
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
			Salary:    salary,
		},
	}
}

func (man *Man) String() string {
	return fmt.Sprintf(
		"My full-name is %v %v. I am %d years old.", man.FirstName,
		man.LastName, man.Age,
	)
}

type Woman struct {
	Human
}

func NewWoman(firstName string, lastName string, age int8, salary int) *Woman {
	return &Woman{
		Human: Human{
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
			Salary:    salary,
		},
	}
}

func (woman *Woman) String() string {
	return fmt.Sprintf(
		"My full-name is %v %v. My salary is %d$.", woman.FirstName,
		woman.LastName,
		woman.Salary,
	)
}

func ExampleStringers() {
	richard := NewMan("Richard", "Richard", 45, 0)
	fmt.Println(richard)

	nicole := NewWoman("Nicole", "Nicole", 40, 50000)
	fmt.Println(nicole)
}
