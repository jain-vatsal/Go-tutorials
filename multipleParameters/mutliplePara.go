package main

import "fmt"

// in case of multiple parameters, it is okay to declare the type at the after the last variable

// x and y are integers and name is a string
func printData(x, y int, name string) {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(name)
}

// if the integers were not in order then they would have to be defined seperately

func main() {
	printData(10, 20, "vatsal")
}
