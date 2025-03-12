package datastructures

import (
	"fmt"
	"reflect"
)

func ExampleMaps() {
	// Creating an empty, unintialized map (nil)
	var map1 map[string]int

	// Creating an empty, intialized map (empty)
	map2 := map[string]int{}
	map3 := make(map[string]int)

	fmt.Println("Looks at these empty maps:")
	fmt.Println(map1, reflect.TypeOf(map1))
	fmt.Println(map2, reflect.TypeOf(map2))
	fmt.Println(map3, reflect.TypeOf(map3))

	fmt.Println()

	// Testing if a key is present in map1
	if _, ok := map1["answer"]; !ok {
		fmt.Println("Shoot! There is no answer in map1, it's unintialized!")
	}

	fmt.Println()

	// Inserting or updating an element in map2
	map2["answer"] = 15
	fmt.Printf("The answer to the problem 1 + 5 must be.. %d?\n", map2["answer"])
	fmt.Println("Ahh, wait! This is not JavaScript..")
	map2["answer"] = 6
	fmt.Printf("The answer to the problem 1 + 5 must be.. %d?\n", map2["answer"])
	fmt.Println("Yay, I am not that bad in math!")
	fmt.Println(map2)
	delete(map2, "answer")
	fmt.Println("Nobody should know the answer! Deleting...")
	fmt.Println(map2)

	fmt.Println()

	// Initializing map1 with some values
	map1 = map[string]int{
		"Ivan":    14,
		"Kelalos": 18,
		"Artyom":  19,
		"Zakhar":  21,
	}

	// Iterating over map1
	for name, age := range map1 {
		fmt.Printf("%s is %d years old.\n", name, age)
	}
}
