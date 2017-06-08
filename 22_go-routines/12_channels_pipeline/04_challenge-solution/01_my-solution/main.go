package main

import "fmt"

func main() {

	in := gen() 	// Gives a receive only channel

	f := factorial(in)	// Factorial takes a receive only channel and also returns a receive only channel

	for n := range f{ 	// Then we range over that receive only channel
		fmt.Println(n)	// and print ou its result.
	}
}

func gen() <- chan int{
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++{
			for j := 3; j < 13; j++{
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <- chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int { 	// Takes each of the numbers passed on by the factorial function
	total := 1	// runs the function bellow and prints it out to "out"
	for i := n; i >0; i -- { 	// in the function above.
		total *= i
	}
	return total
}