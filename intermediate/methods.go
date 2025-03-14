/*
Method - a function associated with a specific type.
*/

package intermediate

import "fmt"

type Person struct {
	Name      string
	Age       uint8
	Interests []string
}

/*
Value Receiver vs Reference Receiver

1) Value Receiver if:
    - the method does not modify the receiver
2) Reference Receiver if:
    - the method modifies the receiver
    - the receiver is large and you want to avoid copying it
*/

// 1) Value Receiver
func (p Person) IntroduceMyself() {
	fmt.Println("Hey, there!")
	fmt.Printf("My name is %s and I am %d years old.\n\n", p.Name, p.Age)

	fmt.Println("A list of my interests:")
	for _, interest := range p.Interests {
		fmt.Printf("- %s\n", interest)
	}
}

// 2) Reference Receiver
func (p *Person) GetOlder() {
	p.Age++
}

func ExampleMethods() {
	me := Person{
		Name:      "Daniil",
		Age:       18,
		Interests: []string{"programming", "sports", "self-development"},
	}
	me.IntroduceMyself()

	fmt.Println()

	me.GetOlder()
	me.IntroduceMyself()

	fmt.Println()

	// Method receivers take either a value or a pointer
	mePtr := &me
	mePtr.GetOlder()
	me.IntroduceMyself()

	fmt.Println()
}
