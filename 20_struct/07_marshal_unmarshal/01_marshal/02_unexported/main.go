package main

import ("fmt"
	"encoding/json"
)

type Person struct {
	first 	string
	last	string
	age 	int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
