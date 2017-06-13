package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0])
}

/*Modify the echo program to also print os.Args[0], the name of the command that invoked it.
 */
