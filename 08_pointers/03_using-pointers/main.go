package main

import "fmt"

func main()  {

	a := 43

	fmt.Println(a) 	//43
	fmt.Println(&a)	//0xc04200c240

	var b *int = &a
	fmt.Println(b)	//0xc04200c240
	fmt.Println(*b)	//43

	*b = 42 //b says: "The value at this address, change it to 42
	fmt.Println(a)	//42
}

//This is useful.
//We can pass a memory address instead of a bunch of values (we can pass a reference)
//and then we can still change the value of whatever is stored at the memory address
//this makes our programs more performant
//we don't have to pass around large amounts of data
//we only have to pass around addresses

//Everything is PASS BY VALUE in Go, BTW.
//When we pass a memory address, we are passing a value