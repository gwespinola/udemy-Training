package main

import "fmt"

func main()  {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar()  {
	for i :=0; i < 1000; i++ {
		fmt.Println("Bar:", i)
	}
}
	// This program gives no input because we have three threads running at the same time and Main exists before returning a result.