package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
		}
		done <- true
	}()

	// We block here until done <- true
	<-done
	<-done
	close(c)

	// To unblock above
	// we need to take values off chan c here
	// but we never get here, because we're blocked above
	for n := range c {
		fmt.Println(n)
	}
}
