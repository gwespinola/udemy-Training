package main

import	"fmt"

func main()  {

	var small_number int32
	var big_number int32

	fmt.Println("Insert small number")
	fmt.Scan(&small_number)
	fmt.Println("Insert larger number")
	fmt.Scan(&big_number)
	fmt.Println("The remainder is ", big_number%small_number)

}