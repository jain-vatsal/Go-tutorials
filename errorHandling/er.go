package main

import (
	"fmt"
	"strconv"
)

// func getUser() (user, error) {

// }

// func getUserProfile() (profile, error) {

// }
func checkIt() {
	// user, err := getUser()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(user)

	// profile, err := getUserProfile(user.ID)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(profile)

	i, err := strconv.Atoi("42b")
	if err != nil {
		fmt.Println("Couldn't convert :", err)
		return
	}
	fmt.Println(i)
}

/*
Built-in Error interface in go

type error interface{
	Error() string
}
*/
