package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// This results in a deadlock.
// Can you know why?
// And what would you to do fix it?
