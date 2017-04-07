package main

import "fmt"

func main()  {
	m := make([]string, 1, 25)
	fmt.Println(m)		//[]
	changeMe(m)
	fmt.Println(m)		//[Gil]
}

func changeMe(z []string) {
	z[0] = "Gil"
	fmt.Println(z)		//[Gil]
}
