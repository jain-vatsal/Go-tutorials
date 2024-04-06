package main

import "fmt"

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	// ?
	disc := calculateDiscount(numMessages)
	baseBill := calculateBaseBill(costPerMessage, numMessages)

	finalBill := baseBill * (1.0 - disc)
	return finalBill
}

func calculateDiscount(messagesSent int) float64 {
	// ?
	if messagesSent > 5000 {
		return 0.2
	} else if messagesSent > 1000 {
		return 0.1
	}
	return 0
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}

func main() {
	bill := calculateFinalBill(100, 10001)
	fmt.Println(bill)
}
