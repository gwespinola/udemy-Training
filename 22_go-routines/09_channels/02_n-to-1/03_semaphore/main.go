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

	go func() {
		<-done // This is how you throw  a value away from the channel.
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
