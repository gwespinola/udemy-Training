package main

import "fmt"

var result int32

func makeInt() {
	x := 10 % 2
	fmt.Println(x)
	if x != 1 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func main() {
	makeInt()
	fmt.Println()
}
