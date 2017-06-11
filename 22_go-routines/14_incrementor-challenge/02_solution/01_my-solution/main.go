package main

import (
	"fmt"
	"time"
	"math/rand"
)

var count int64


func main() {
	c:= fanIn(incrementor("1"), incrementor("2"))

	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Final counter: ", count)
}

func incrementor(s string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 20; i++ {
			c <- fmt.Sprintln("Process: "+s+" printing:", i)
			time.Sleep(time.Duration(rand.Intn(1)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <- input1
			count++
		}
	}()
	go func() {
		for {
			c <- <-input2
			count++
		}
	}()
	return c
}


/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/