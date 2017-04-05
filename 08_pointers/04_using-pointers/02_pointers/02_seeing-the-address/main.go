package main

import "fmt"

func zero(x *int) {
	fmt.Println(x)
	*x = 0
}

func main()  {
	x := 0
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)	//x is 0
}
