package main

import "fmt"

func main() {
	c := incrementor()
	for n := range puller(c) { // This code is differentiated from the one in the previous exercise, because here we used regular code substitution
		fmt.Println(n) // and instead of having a variable "cSum" assigning the channel result to the expression range we can use the principle
	} // of equality and substitute the variable "cSum" by the result of channel "puller(c)".
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

/* Here we have some function that is returning a channel and instead of repeating the code we are just ranging
over that
*/
