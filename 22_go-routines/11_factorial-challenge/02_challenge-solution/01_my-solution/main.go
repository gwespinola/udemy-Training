package main

import (
	"fmt"
)

func main() {
	f := factorial(4)
	for n := range f {
		fmt.Println("Total:", n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	var total int = 1
	go func() {
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}


