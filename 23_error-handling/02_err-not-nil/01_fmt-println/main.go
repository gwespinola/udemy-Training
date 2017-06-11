package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("Err happened", err)
	//	log.Println("Err happened", err)
	//	log.Fatalln(err)
	//	panic(err)
	}
}

// Println formats using the default formats for its operands and writes to standard output.