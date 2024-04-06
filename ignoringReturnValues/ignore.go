package main

import "fmt"

func returnMultipleValues(x, y int) (int, int) {
	return x, y
}
func main() {
	x, _ := returnMultipleValues(30, 4)
	fmt.Println(x)
}
