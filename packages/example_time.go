package packages

import (
	"fmt"
	"time"
)

func ExampleTime() {
	// time.Now - get the current time.
	now := time.Now()
	fmt.Println("Current time:", now)

	// time.Format - format time using a custom layout.
	// 2006-01-02 15:04:05 - is a hard-coded reference date in Go formatting system.
	// time.Sleep - pauses for specific period of time
	// time.Second - returns a constant representing a duration of one second
	formattedNow1 := now.Format("2006-01-02 15:04:05")
	formattedNow2 := now.Format("01-02-2006 15:04:05")
	fmt.Println("Formatted current time 1:", formattedNow1)
	time.Sleep(time.Second * 2)
	fmt.Println("Formatted current time 2:", formattedNow2)
}
