package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 32, 3}

	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)
}
