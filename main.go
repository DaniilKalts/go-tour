/*
Package: is a collection of Go files (typically in the same directory).
*/

// A package name for runnable programs must be "main"
package main

import (
	"go-tour/basics/datastructures"
	"go-tour/utils"
)

func main() {
	utils.PrintSyntaxSection("Arrays", datastructures.ExampleArrays)
}
