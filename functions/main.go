package main

import "fmt"

// types come after
func concat(s1 string, s2 string) string {
	return s1 + s2
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

func main() {
	test("Vatsal", "Jain")
	test("Vatsal", " Jain")
	test("Vatsal", " Sethi")
}
