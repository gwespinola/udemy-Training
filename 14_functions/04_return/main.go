package main

import "fmt"

func main()  {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(fname, lname string) string  {
	return fmt.Sprint(fname, lname)
}

//Sprint function is a string print.
//It will do the printing but to a string, not to a standard out i.e. monitor