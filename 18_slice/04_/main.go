package main

import "fmt"

func main()  {

	mySlice := []int{1, 3, 5, 7, 9, 11,}	//the problem here is that we have not set the capacity of the underlying array,
						// so the present expression will have length 6 and a capacity of 6
	for i, value := range mySlice  {
		fmt.Println(i, " - ", value)
	}
}
