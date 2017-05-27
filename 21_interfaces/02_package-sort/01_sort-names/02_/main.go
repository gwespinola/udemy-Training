package main

import (
	"fmt"
	"sort"
)

type people []string

func strings(z people) {
	return sort.Strings(people)
}

func main()  {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	fmt.Println(strings(people)
}