package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann")) // boring joe and boring ann both return a channel and fanIn takes two inputs
	for i := 0; i < 10; i++ {                // of the receive only channel and the channel received has strings on it.
		fmt.Println(<-c) // This channel prevents the program from exiting.
	}
	fmt.Println("You are both boring; I'm leaving.")
}

func boring(msg string) <-chan string { // This function takes a string and then it makes a channel.
	c := make(chan string)
	go func() { // launches a go function and then returns the channel
		for i := 0; ; i++ { // then it writes a string to the channel
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // and that gets returned so we return a channel
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string { // This function takes two channels: input 1 and input2.
	c := make(chan string)
	go func() {
		for {
			c <- <-input1 // We take the value of input1 and then puts it on c which
		} // is a brand new channel.
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25
*/

/*
FAN OUT
Multiple functions reading from the same channel until that channel is closed
FAN IN
A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.
PATTERN
there's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.
source:
https://blog.golang.org/pipelines
*/
