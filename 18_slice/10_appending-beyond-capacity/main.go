package main

import "fmt"

func main()  {

	greeting := make([]string, 3, 5)
	//3 is length - number of elements in the array referred by the slice
	//5 is capacity - number of elements in the underlying array
	greeting[0] = "Good Morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "suprabadham")
	greeting = append(greeting, "Zǎo'ān")
	greeting = append(greeting, "Ohayou gozaimasu")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))

}
//To add the 4th element we need to use the append command and golang automatically doubles the size of the array.
