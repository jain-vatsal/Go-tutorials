package main

import "fmt"

type car struct {
	maker string
	model string
}

type truck struct {
	car     // truck is automatically assigned the attributes of car
	bedsize int
}

func main() {
	lanesTruck := truck{
		bedsize: 10,
		car: car{
			maker: "toyota",
			model: "camry",
		},
	}

	fmt.Println(lanesTruck.bedsize)
	fmt.Println(lanesTruck.maker) //  no need for car.maker
	fmt.Println(lanesTruck.model)
}
