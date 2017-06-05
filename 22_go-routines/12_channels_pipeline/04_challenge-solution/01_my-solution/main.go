package main

import "fmt"

func main() {
	for n := range fac(gen(100)) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for nums := 0; nums > 0; nums-- {
			out <- nums
		}
		close(out)
	}()
	return out
}

func fac(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
