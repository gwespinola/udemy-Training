package main

import "fmt"

type Person struct {
	First	string
	Last 	string
	Age 	int
}

type  doubleZero struct {
	Person
	First		string
	LicenseToKill	bool
}

func main() {
	p1 := doubleZero{
		Person: Person{
			First: "James",
			Last: "Bond",
			Age: 20,
		},
		First:		"Double Zero Seven",
		LicenseToKill: 	true,
	}
	p2 := doubleZero{
		Person: Person{
			First: "Miss",
			Last: "Moneypenny",
			Age: 19,
		},
		First:		"If looks could kill",
		LicenseToKill:	false,
	}
//	Fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(p1.First, p1.Person.First)
	fmt.Println(p2.First, p2.Person.First)
}
