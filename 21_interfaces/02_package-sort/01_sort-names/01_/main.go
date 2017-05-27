package main

import "fmt"

type people []string

func main() {
	var studyGroup = people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
}
