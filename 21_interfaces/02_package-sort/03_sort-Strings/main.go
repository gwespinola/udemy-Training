package main

import (
	"fmt"
	"sort"
)

func main() {
	studyGroup := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println(studyGroup)
}
