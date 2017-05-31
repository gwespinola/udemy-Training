package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println(studyGroup)
}
