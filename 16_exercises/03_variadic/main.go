package main

import "fmt"

func main() {
	x := []int{1, 5, 9, 18, 29, 34}

	var number int = x[0]

	for i := 0; i < len(x); i++ {
		if x[i] > number {
			number = x[i]
		}
	}
	fmt.Println("Biggest number is: ", number)
}
