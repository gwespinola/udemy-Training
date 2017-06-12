package main

import (
	"fmt"
	"os"
)
// The variables declarations declares two variables of type string. A variable can be initialised as part of its declaration.
// If it's not explicitly initialised, it is implicitly initialised to the zero type of its value, which is 0 for numeric types
// and the empty string " " for strings.
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
}
