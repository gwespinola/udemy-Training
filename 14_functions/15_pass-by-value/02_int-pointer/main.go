package main

import "fmt"

func main()  {

	age := 44
	fmt.Println(&age) //0xc04203a1d0

	changeMe(&age)

	fmt.Println(&age)//0xc04203a1d0
	fmt.Println(age)//24
}

func changeMe(z * int)  {
	fmt.Println(z)	//0xc04203a1d0
	fmt.Println(*z)	//44
	* z = 44
	fmt.Println(z)	//0xc04203a1d0
	fmt.Println(*z)	//24
}
