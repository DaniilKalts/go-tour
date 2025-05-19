/*
Package: is a collection of Go files (typically in the same directory).
*/

// A package name for runnable programs must be "main"
package main

import (
	"github.com/DaniilKalts/go-tour/intermediate"
	"github.com/DaniilKalts/go-tour/utils"
)

func main() {
	utils.PrintSyntaxSection("Errors", intermediate.ExampleErrors)
}
