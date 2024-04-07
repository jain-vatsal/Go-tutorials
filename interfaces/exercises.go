package main

import "fmt"

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	name    string
	message string
}

type sendSomething struct {
	phoneNumber int
	message     string
}

func (b birthdayMessage) getMessage() string {
	return fmt.Sprintf("Your message is " + b.message)
}

func (s sendSomething) getMessage() string {
	return fmt.Sprintf("Your message is " + s.message)
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

func main() {
	vat := message(birthdayMessage{
		name:    "Vatsal Jain",
		message: "Kaise ho",
	})

	sendMessage(vat)

	vat1 := message(sendSomething{
		phoneNumber: 12345678910,
		message:     "How are you",
	})

	sendMessage(vat1)
}
