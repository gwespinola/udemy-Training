package main

import "fmt"

 // This function is supposed to return an err due to a lack of interface, so we cannot use assertion.
func main() {
	name := "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}