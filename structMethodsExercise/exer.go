package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	ans := "Authorization: Basic " + a.username + ":" + a.password
	return ans
}

func test(a authenticationInfo) {
	fmt.Println(a.getBasicAuth())
}
func main() {
	vat := authenticationInfo{
		username: "vatsalJain",
		password: "1234",
	}
	test(vat)
}
