package main

import "fmt"

type Person struct {
	Name 	string
	Age	int
}

type doubleZero struct {
	Person
	LicenseToKill	bool
}

func (p Person) Greeting()  {
	fmt.Println("I'm just a regular person")
}

func (dz doubleZero) Greeting()  {
	fmt.Println("Miss MoneyPenny, so good to see you.")
}

func main()  {
	p1 := Person{
		Name: "Ian Flemming",
		Age: 44,
	}

	p2 := doubleZero{
		Person: Person{
			Name: "James Bond",
			Age: 23,
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
}
