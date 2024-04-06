package main

import "fmt"

func main() {

	// %v - interpolates the default representation
	fmt.Printf("I am %v years old", 21)
	fmt.Printf("I am %v years old", "way too many")

	// %s - interpolates a string
	fmt.Printf("I am %s years old", "not too many")

	s := fmt.Sprintf("I am %f years old", 10.523)
	// I am 10.523000 years old
	fmt.Printf(s)

	// The ".2" rounds the number to 2 decimal places
	r := fmt.Sprintf("I am %.2f years old", 10.523)
	fmt.Printf(r)
}
