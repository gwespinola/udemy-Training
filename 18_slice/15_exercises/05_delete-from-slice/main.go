package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 5)
	mySlice[0] = 48
	mySlice[1] = 8
	mySlice[2] = 64
	mySlice[3] = 94
	mySlice[4] = 72
	fmt.Println(mySlice)
	mySlice = append(mySlice[:1], mySlice[3:]...)
	fmt.Println(mySlice)
	fmt.Println(cap(mySlice))
}

//delete an element from a slice
