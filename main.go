/*
Package: is a collection of Go files (typically in the same directory).
*/

// A package name for runnable programs must be "main"
package main

import (
	"go-tour/packages"
	"go-tour/utils"
)

func main() {
	utils.PrintSection("Fmt", packages.ExampleFmt)
	utils.PrintSection("Math", packages.ExampleMath)
	utils.PrintSection("Time", packages.ExampleTime)
}
