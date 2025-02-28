// By convention, the package name is the same as the last element of the import path.
// For instance, the "math/rand" package comprises files that begin with the statement package rand.

package packages

import (
	"fmt"
	"math/rand"
)

func ExampleMath() {
	// rand.Intn - returns a random integer from 0 to n - 1
	fmt.Println("Just a random number from 1 to 10:", rand.Intn(10)+1)
}
