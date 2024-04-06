package main

import "fmt"

func main() {
	const name = "Vatsal Jain"
	const openRate = 101.5

	msg := fmt.Sprintf(
		"Hi %s, your open rate is %.1f percent\n",
		name,
		openRate,
	)

	fmt.Println(msg)
}
