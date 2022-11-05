package main

import "fmt"

func main() {
	// var colors map[string]string // or
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#443953",
		"white": "#fffff",
	}
	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
