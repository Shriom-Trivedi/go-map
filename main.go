package main

import "fmt"

func main() {
	colors:= map[string]string {
		"red": "ff0000",
		"blue": "#00FFFF",
		"green": "#088F8F",
	}

	// var colors map[string]string
	// colors:= make(map[int]string)
	// colors[10] = "red"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}