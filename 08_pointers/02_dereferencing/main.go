package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)	//43
	fmt.Println(&a)	//0xc04200c240

	var b *int = &a
	fmt.Println(b)	//0xc04200c240
	fmt.Println(*b)	//0xc04200c240
}
//	b is an int pointer
//	b points to the memory address where an int is stored
//	to see the value of the memory address, add a * in front of
//	this is know as dereferencing
//	the * is an operator in this case
