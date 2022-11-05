package main

import "fmt"

type location struct {
	longitude float64
	latitude  float64
}

func main() {
	newYork := location{
		latitude:  40.73,
		longitude: -73.93,
	}

	newYork.changeLatitude()

	fmt.Println(newYork)

	name := "Bill"
	updateValue(name)
	fmt.Println(*&name)

	namePointer := &name
	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func (lo *location) changeLatitude() {
	(*lo).latitude = 41.0
}

func updateValue(n string) {
	n = "Alex"
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
