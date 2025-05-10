/*
Generic(type parameters) - a language feature, allowing to work with many data types.
*/

package intermediate

import "fmt"

// Example 1:

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func sumNumbers[T Number](numbers []T) T {
	var total T

	for i := range numbers {
		total += numbers[i]
	}

	return total
}

// Example 2:

func indexAt[S ~[]E, E comparable](slice S, element E) int {
	for i, value := range slice {
		if element == value {
			return i
		}
	}

	return -1
}

// Example 3:

type element[T any] struct {
	next *element[T]
	val  T
}

type LinkedList[T any] struct {
	head *element[T]
}

func (lst *LinkedList[T]) Print() {
	for curr := lst.head; curr != nil; curr = curr.next {
		fmt.Println(curr.val)
	}
}

func (lst *LinkedList[T]) Push(val T) {
	if lst.head == nil {
		lst.head = &element[T]{val: val}
		return
	}

	curr := lst.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &element[T]{val: val, next: nil}
}

// Example 4

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

func (set *Set[T]) Add(val T) {
	set.m[val] = struct{}{}
}

func (set *Set[T]) Remove(val T) {
	delete(set.m, val)
}

func (set *Set[T]) Contains(val T) bool {
	_, ok := set.m[val]
	return ok
}

func (set *Set[T]) Len() int {
	return len(set.m)
}

func ExampleGenerics() {
	// Example 1:

	decimals := []int{1, 2, 3, 4, 5}
	floats := []float32{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Printf("Sum of %v: %v", decimals, sumNumbers(decimals))
	fmt.Println()

	fmt.Printf("Sum of %v: %v", floats, sumNumbers(floats))
	fmt.Println()
	fmt.Println()

	// Example 2:

	fmt.Printf("Index of %v in %v: %d\n", 3, decimals, indexAt(decimals, 3))
	fmt.Printf("Index of %v in %v: %d\n", 6, decimals, indexAt(decimals, 6))
	fmt.Printf("Index of %v in %v: %d\n", 4.4, floats, indexAt(floats, 4.4))
	fmt.Printf("Index of %v in %v: %d\n", 1.0, floats, indexAt(floats, 1.0))
	fmt.Println()

	// Example 3:

	lst := LinkedList[int]{}
	lst.Push(1)
	lst.Push(2)
	lst.Push(3)
	lst.Print()
	fmt.Println()

	// Example 4:
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(2)

	fmt.Printf("Contains 2? %v\n", set.Contains(2))
	fmt.Printf("Contains 3? %v\n", set.Contains(3))
	fmt.Printf("Length: %d\n", set.Len())
}
