package main

import "fmt"

func main()  {

	customerNumber := make([]int, 3)
	//3 is length and capacity
	//length - number of elements referred to by the slice
	//capacity - number of elements in the underlying array
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	//3 is length - number of elements referred to by the slice
	//5 is capacity - number of elements in the underlying array
	//you could also do it like this:

	greeting[0] = "Good Morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"

	fmt.Println(greeting[2])
}

/*If we think our slice might grow, we can set a capacity larger than length. THIS GIVES OUR SLICE ROOM TO GROW WITHOUT GOLANG HAVING TO CREATE A NEW UNDERLYING ARRAY EVERY TIME OUR SLICE GROWS. When the slice exceeds capacity, then a new underlying array will be created. These arrays double in size each time they are created (2, 4, 8, 16) up to a certain point, and then they scale in some smaller proportion. */
