package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	}
	fmt.Println(s)
	fmt.Println(sep)
}

/*Modify the echo program to print the index and value of each of its arguments, one per line.
 */
