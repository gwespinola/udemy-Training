package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)
	fmt.Println(mySlice)
}

//create a slice of ints using make; set the length to 5 and capacity to 10
