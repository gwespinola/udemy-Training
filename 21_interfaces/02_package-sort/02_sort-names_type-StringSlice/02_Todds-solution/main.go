package main

import (
	"fmt"
	"sort"
)

//This parte of the code is just to show that you can attach a method to a variable
//type t []string
//
//var s = []string{"Zeno", "John", "Al", "Jenny"}
//
//func (z t) something() {
//	fmt.Println(z)
//}

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.StringSlice(s).Sort()
	//	The code bellow  is an alternative way of writing the sort function.
	//	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
