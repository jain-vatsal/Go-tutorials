// if there is no need of more than one instance of a struct, it can be simply assigned to a variable

package main

import "fmt"

// this is a nested anonymous struct
type car struct {
	Make   string
	Model  string
	Height int
	Width  int
	Wheel  struct {
		Radius   int
		Material string
	}
}

func main() {
	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "Audi",
		Model: "Q3",
	}

	fmt.Println(myCar.Make)
	fmt.Println(myCar.Model)

	newCar := car{
		Make:   "XUV",
		Model:  "500",
		Height: 3,
		Width:  3,
		Wheel: struct {
			Radius   int
			Material string
		}{
			Radius:   1,
			Material: "Carbon",
		},
	}

	// printing a struct
	fmt.Printf("%+v\n", newCar)
}
