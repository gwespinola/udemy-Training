package main

import "fmt"

func main() {
	mySlice := make([]string, 1)
	mySlice[0] = "Chris"
	mySlice[1] = "Johann"
	fmt.Println(mySlice)
}

//create a slice then make your program throw an “index out of range” error
