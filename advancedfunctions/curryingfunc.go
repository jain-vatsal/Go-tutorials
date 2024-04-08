package main

import "fmt"

// import (
// 	"fmt"
// )

// func main() {
// 	squareFunc := selfMath(multiply)
// 	doubleFunc := selfMath(addition)

// 	fmt.Println(squareFunc(5))
// 	fmt.Println(doubleFunc(5))

// }

// func multiply(x, y int) int {
// 	return x * y
// }

// func addition(x, y int) int {
// 	return x + y
// }

// func selfMath(mathFunc func(int, int) int) func(x int) int {
// 	return func(x int) int {
// 		return mathFunc(x, x)
// 	}
// }

/*

The Mailio API needs a very robust error-logging system so we can see when things are going awry in the back-end system. We need a function that can create a custom "logger" (a function that prints to the console) given a specific formatter.

Complete the getLogger function. It should return a new function that prints the formatted inputs using the given formatter function. The inputs should be passed into the formatter function in the order they are given to the logger function.

*/

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	// ?
	return func(a, b string) {
		fmt.Println(formatter(a, b))
	}
}
