package main

import "fmt"

func fibonacci(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main()  {
	n := uint(3999999)
	fmt.Println(fibonacci(n))

}