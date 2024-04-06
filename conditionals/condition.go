package main

import "fmt"

func main() {
	height := 7

	if height > 6 {
		fmt.Println("You are super tall!")
	} else if height > 4 {
		fmt.Println("You are tall enough!")
	} else {
		fmt.Println("You are short!")
	}
}
