package main

import "fmt"

func main() {
	messageFromDoris := []string{
		"Hi There",
		"How are you",
		"So cool",
	}

	numMessages := float64(len(messageFromDoris))
	costPerMessage := 0.02

	totalCost := costPerMessage * numMessages

	fmt.Printf("Doris spent %0.2f on text messages on this line\n", totalCost)
}
