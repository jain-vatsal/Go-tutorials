package main

import "fmt"

// in go, variables are passed by value
func increment(x int) int {
	x++
	return x
}
func main() {
	x := 5
	fmt.Println(increment(x))

	// x is unchanged
	fmt.Println(x)
}
