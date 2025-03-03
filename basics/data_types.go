package basics

/* BOOLEAN */
var (
	isGopher bool // True or False
)

/* NUMERIC */
var (
	// integers
	a int8  // int8 range: -128 to 127
	b int16 // int16 range: -32768 to 32767
	c int32 // int32 range: -2147483648 to 2147483647
	d int64 // int64 range: -9223372036854775808 to 9223372036854775807
	e int   // int range: platform-dependent (-2^31 to 2^31-1 on 32-bit, -2^63 to 2^63-1 on 64-bit)

	// unsigned integers
	f byte   // byte (alias for uint8) range: 0 to 255
	g uint16 // uint16 range: 0 to 65535
	h uint32 // uint32 range: 0 to 4294967295
	i uint64 // uint64 range: 0 to 18446744073709551615
	j uint   // uint range: platform-dependent (0 to 2^32-1 on 32-bit, 0 to 2^64-1 on 64-bit)

	// floating-point numbers
	k float32 // float32 range: approx. 1.4e-45 to 3.4e38
	l float64 // float64 range: approx. 5e-324 to 1.8e308

	// complex numbers
	m complex64  // complex64: each part is float32 (approx. 1.4e-45 to 3.4e38)
	n complex128 // complex128: each part is float64 (approx. 5e-324 to 1.8e308)
)

/* STRINGS */
var (
	greeting string // any sequence of characters
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
