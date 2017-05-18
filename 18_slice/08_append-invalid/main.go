package main

import "fmt"

func main()  {

	greeting := make([]string, 3, 5)
	//3 is length - number of elements in the array referred by the slice
	//5 is capacity - number of elements in the underlying array
	greeting[0] = "Good Morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting[3] = "suprabadham"

	fmt.Println(greeting[2])

}
//This is program is supposed to return an error because we inserted a fourth element while the length is limited to 3. The 4th element is outside the range and therefore returns and error.
