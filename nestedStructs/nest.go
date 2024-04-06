package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(m messageToSend) bool {
	if m.sender.name == "" || m.sender.number == 0 {
		return false
	}
	if m.recipient.name == "" || m.recipient.number == 0 {
		return false
	}
	return true
}
func main() {

}
