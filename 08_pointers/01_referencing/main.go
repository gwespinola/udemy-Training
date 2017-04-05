package main

import "fmt"

func main()  {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b * int = &a

	fmt.Println(b)
}

//The above code makes b a pointer to the memory address where an int is stored
//b is of tyoe "int pointer"
//*int -- the * is part of the type -- b is of part *int
