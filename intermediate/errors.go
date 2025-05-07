/*
error - the interface for errors.
*/
package intermediate

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	const layout = "2006-01-02 15:04:05"
	return fmt.Sprintf("%s: %s", e.When.Format(layout), e.What)
}

func GetPension(age int) (int, *MyError) {
	if age < 65 {
		return 0, &MyError{
			When: time.Now(),
			What: fmt.Sprintf("You are too young! Wait %d years.", 65-age),
		}
	}
	return 150, nil
}

func ExampleErrors() {
	if _, err := GetPension(18); err != nil {
		fmt.Println(err)
	}
	if _, err := GetPension(65); err != nil {
		fmt.Println(err)
	}
}
