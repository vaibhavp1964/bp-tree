package utils

import "fmt"

// AssertEquals ...
func AssertEquals[T comparable](a, b T) {
	if a != b {
		panic(fmt.Sprintln(a, "and", b, "must be equal"))
	}
}
