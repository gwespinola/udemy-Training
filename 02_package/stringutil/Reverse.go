// Package stringutil contains utility functions for working with
package stringutil

//Reverse returns its arguments string reversed rune-wise left to right
func Reverse(s string) string {
	return reverseTwo(s)
}