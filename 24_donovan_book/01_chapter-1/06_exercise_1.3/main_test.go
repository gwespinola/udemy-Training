package main

import (
	"testing"
	"os"
	"fmt"
	"strings"
)

func BechmarkJoin(b *testing.B) {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}



/*Experiment to measure the difference in running time between our potentially inefficient versions and the ones that
 use string.Join
*/