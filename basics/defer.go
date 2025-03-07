/*
Defer - is a statement that schedules a function call to be executed after the surrounding function completes.
*/

package basics

import (
	"fmt"
	"os"
	"time"
)

func ExampleDefer() {
	// Get the current working directory
	dir, _ := os.Getwd()

	file, err := os.Open(dir + "/data/example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Ensure that the file is closed after the program finishes
	defer file.Close()

	fmt.Println("File opened successfully")
	fmt.Println()
	fmt.Println("Now I'll print numbers from 10 to 1:")

	for i := 1; i < 11; i++ {
		defer time.Sleep(time.Millisecond * 500)
		defer fmt.Println(i)
	}
}
