package main

import "fmt"

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(0, costPerSend, numLastMonth)
	thisMonthBill := getBillForMonth(0, costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(bill, costPerSend, messagesSent int) int {
	bill = costPerSend * messagesSent
	return bill
}

func main() {
	ans := monthlyBillIncrease(10, 10, 20)
	fmt.Println(ans)
}
