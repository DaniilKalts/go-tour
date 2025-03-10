/*
Array - a fixed-size, contiguous collection of elements of a single type.
*/
package datastructures

import "fmt"

func ExampleArrays() {
	var ipAddresses [3]string = [3]string{"17.172.224.47", "123.89.46.72", "192.168.43.241"}
	fmt.Printf("Printing random IP addresses using %%v : %v\n", ipAddresses)
	fmt.Printf("Printing random IP addresses using %%#v: %#v\n", ipAddresses)
}
