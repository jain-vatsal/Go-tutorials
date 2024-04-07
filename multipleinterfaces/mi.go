// ASSIGNMENT
// Complete the required methods so that the email type implements both the expense and formatter interfaces.

// COST()
// If the email is not "subscribed", then the cost is 5 cents for each character in the body. If it is, then the cost is 2 cents per character.

// Return the total cost of the entire email in cents.

// FORMAT()
// The format method should return a string in this format:

// 'CONTENT' | Subscribed
//
// If the email is not subscribed, change the second part to "Not Subscribed":

// 'CONTENT' | Not Subscribed
//
// The single quotes are included in the string, and CONTENT is the email's body. For example:

// 'Hello, World!' | Subscribed

package main

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	}
	return 0.02 * float64(len(e.body))
}

func (e email) format() string {
	ans := "'" + e.body + "'" + " | "
	if e.isSubscribed {
		ans = ans + "Subscribed"
	} else {
		ans = ans + "Unubscribed"
	}
	return ans
}

type expense interface {
	cost() float64
}

type formater interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
