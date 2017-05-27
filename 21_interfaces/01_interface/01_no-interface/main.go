package main

import "fmt"

// Creates a user-defined type and call it a square
type Square struct {
	side float64
}

// We then attach a method to this type we just created and if we wanted we could that method right out of the type
func (z Square) area() float64 {
	return z.side * z.side
}

// Here we create a new user-defined type called Shape. The type Shape is of type interface.
// Interfaces define functionality and so anything that has this method signature "area() float64" implements the Shape interface.
// Since Square has a method that implements the Shape interface, Square now implements the Shape interface
type Shape interface {
	area() float64
}

// Because Square implements the Shape interface we can now create a new function and it can take a type. The type this function has taken is the type Shape.
// So you can pass anything that implements type Shape into this function.
func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

// So right bellow we calling the function "info" and are passing in "s". And "s" is a variable of type Square.
// But since Square implements the Shape interface and implements it by having the Square method we can then pass it as info because Square is a Shape.
func main() {
	s := Square{10}
	info(s)
}
