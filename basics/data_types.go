package basics

import (
	"fmt"
	"reflect"
	"strconv"
)

/* BOOLEAN */
var (
	isGopher bool // True or False
)

/* NUMERIC */
var (
	// integers
	age int8  // int8 range: -128 to 127
	b   int16 // int16 range: -32768 to 32767
	c   int32 // int32 range: -2147483648 to 2147483647
	d   int64 // int64 range: -9223372036854775808 to 9223372036854775807
	e   int   // int range: platform-dependent (-2^31 to 2^31-1 on 32-bit, -2^63 to 2^63-1 on 64-bit)

	// unsigned integers
	ABSOLUTE_PI byte   // byte (alias for uint8) range: 0 to 255
	g           uint16 // uint16 range: 0 to 65535
	h           uint32 // uint32 range: 0 to 4294967295
	i           uint64 // uint64 range: 0 to 18446744073709551615
	j           uint   // uint range: platform-dependent (0 to 2^32-1 on 32-bit, 0 to 2^64-1 on 64-bit)

	// floating-point numbers
	k  float32 // float32 range: approx. 1.4e-45 to 3.4e38
	PI float64 // float64 range: approx. 5e-324 to 1.8e308

	// complex numbers
	m complex64  // complex64: each part is float32 (approx. 1.4e-45 to 3.4e38)
	n complex128 // complex128: each part is float64 (approx. 5e-324 to 1.8e308)
)

/* STRINGS */
var (
	queryId string // any sequence of characters
)

/* AGGREGATE */
var (
	// array
	arr [5]int // fixed-size sequence

	// slice
	slice []int // non-fixed-size sequence

	// maps
	hashMap map[string]int
)

// pointers
var (
	intPtr    *int
	stringPtr *string
)

// structs
type User struct {
	Name string
	Age  int
}

// interfaces
type Student interface {
	argue(teacher, argument string) string
	getMarks() []int8
}

func ExampleDataType() {
	age = 18

	/* GET THE DATA TYPE */
	fmt.Printf("I am %d years old.\n", age)
	fmt.Println("You may be wondering: what type did I use for storing my age..")
	fmt.Println()
	fmt.Println("Well, there are 2 ways (that I am familiar with):")
	fmt.Printf("1) Use format specifier '%%T': I used %T for storing my age in the variable.\n", age)
	fmt.Printf("2) Use reflect package      : I used %v for storing my age in the variable.\n", reflect.TypeOf(age))
	fmt.Println()

	/* TYPE CONVERSION */

	// 1) Direct Type Conversion: T(v)
	PI = 3.141592653589793238462643383279502884197
	fmt.Printf("PI used to be %f with %v data type.", PI, reflect.TypeOf(PI))
	fmt.Println()
	ABSOLUTE_PI = byte(PI)
	fmt.Printf("But now it has turned to %d with %v data type", ABSOLUTE_PI, reflect.TypeOf(ABSOLUTE_PI))
	fmt.Println()
	fmt.Println()

	// 2) strconv package: strconv.Atoi (ASCII to Integer) or strconv.Itoa (Integer to ASCII)
	queryId = "666"

	id, err := strconv.Atoi(queryId)
	if err != nil {
		fmt.Printf("Error converting query ID: %v\n", err)
		return
	}
	fmt.Printf("Boom! We've just converted a query id from %v to %v", reflect.TypeOf(queryId), reflect.TypeOf(id))
	fmt.Println()
}
