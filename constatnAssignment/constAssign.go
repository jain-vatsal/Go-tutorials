package main

import "fmt"

func main() {
	const secInMin = 60
	const min = 60
	const totalSeconds = secInMin * min // this is valid in Go

	fmt.Println(totalSeconds)
}
