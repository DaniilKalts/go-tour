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
	utils.PrintLibrarySection("Image", packages.ExampleImage)
}
