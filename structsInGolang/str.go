package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Println("Sending message : ", m.message, m.phoneNumber)
}
func main() {
	test(messageToSend{
		phoneNumber: 12345678910,
		message:     "Hello there",
	})

	test(messageToSend{
		phoneNumber: 12345610,
		message:     "Hello",
	})
}
