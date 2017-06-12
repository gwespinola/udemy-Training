package main

import (
	"testing"
	"os"
	"fmt"
)

func BechmarkJoin(b *testing.B) {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
}


/*Experiment to measure the difference in running time between our potentially inefficient versions and the ones that
 use string.Join
*/