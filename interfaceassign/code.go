package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type sde struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (s sde) getName() string {
	return s.name
}

func (s sde) getSalary() int {
	return s.hourlyPay * s.hoursPerYear
}

type marki struct {
	name   string
	salary int
}

func (m marki) getName() string {
	return m.name
}

func (m marki) getSalary() int {
	return m.salary
}

func test(e employee) {
	fmt.Println(e.getName())
	fmt.Println(e.getSalary())
}

func main() {
	sde1 := employee(sde{
		name:         "Vatsal Jain",
		hourlyPay:    1000,
		hoursPerYear: 10000,
	})
	test(sde1)

	marki1 := employee(marki{name: "Vatsu", salary: 1000000})
	test(marki1)
}
