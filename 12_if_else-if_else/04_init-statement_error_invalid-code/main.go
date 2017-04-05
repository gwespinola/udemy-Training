package main

import "fmt"

func main()  {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	fmt.Println(food)
}

	//this is just to show that you can't get the scope out of the
	//if statement