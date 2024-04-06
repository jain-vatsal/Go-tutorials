package main

import "fmt"

func main() {
	var username string = "vatsal"
	var password int = 123

	// since go is statically typed, it does not allow concatentation of strings with integers, which is why you need comma as shown below
	fmt.Println("Authorization: Basic", username, ":", password)

	// strings can be concatenated
	var user string = "vatsaljain"

	fmt.Println(username + ":" + user)
}
