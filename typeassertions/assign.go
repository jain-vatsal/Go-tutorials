// type shape interface {
// 	area() float64
// }

// type circle struct {
// 	radius float64
// }

// c, ok := s.(circle)
// if !ok {
// 	// log an error if s isn't a circle
// 	log.Fatal("s is not a circle")
// }

// radius := c.radius

// ASSIGNMENT
// Implement the getExpenseReport function.

// If the expense is an email then it should return the email's toAddress and the cost of the email.
// If the expense is an sms then it should return the sms's toPhoneNumber and its cost.
// If the expense has any other underlying type, just return an empty string and 0.0 for the cost.

package main

func getExpenseReport(e expense) (string, float64) {
	// ?

	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}
	s, ok := e.(sms)
	if ok {
		return s.toPhoneNumber, s.cost()
	}
	return "", 0.0
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {

}
