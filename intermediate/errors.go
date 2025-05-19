/*
error - the interface for errors.

panic - built-in function that triggers an abnormal termination.
recover - built-in function that captures panic and lets program resume normal execution.
*/
package intermediate

import (
	"errors"
	"fmt"
	"time"
)

// ==================

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

// ==================

func SafeGetInt(s []int, i int) (val int, err error) {
	defer func() {
		if r := recover(); r != nil {
			val = -1
			err = errors.New("index out of range")
		}
	}()

	if 0 <= i && i < len(s) {
		return s[i], nil
	}
	panic("index out of range")
}

// ==================

func ExampleErrors() {
	if _, err := GetPension(18); err != nil {
		fmt.Println(err)
	}
	if _, err := GetPension(65); err != nil {
		fmt.Println(err)
	}

	data := []int{10, 20, 30}

	if v, err := SafeGetInt(data, 1); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("data[1] =", v)
	}

	if v, err := SafeGetInt(data, 5); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("data[5] =", v)
	}
}
