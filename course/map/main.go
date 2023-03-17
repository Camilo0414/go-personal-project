package main

import "fmt"

func main() {
	// First approach of how create a map
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"white": "#FFFFFF",
	}

	// Second approach of how create a map
	// var colors map[string]string

	// Third approach of how create a map
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
