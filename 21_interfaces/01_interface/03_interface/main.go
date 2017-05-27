package main

// On this example we have two types of Shape.
import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

// another shape
type circle struct {
	radius float64
}

// Now both Circle and Square implement the Shape interface.
// In Go interfaces are implemented implicitly so you don't have to say anywhere that this implements the interface.
// Jus by having the method with the correct signature that matches how an interface is declared than whatever type has that method that type now implements that interface.
type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

// which implements the shape interface
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Since "info" takes a Shape it can now take a Circle or a Square.
// The signatures have to be the same. The same method and number of parameters and arguments both have to match what was declared wen you declared that interface.
func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

// Here the function runs a different functionality to whatever Type is implementing that.
func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
}
