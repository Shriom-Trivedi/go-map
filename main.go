package main

import "fmt"

func main() {
	// colors:= map[string]string {
	// 	"red": "ff0000",
	// 	"blue": "#00FFFF",
	// 	"green": "#088F8F",
	// }

	// var colors map[string]string
	colors:= make(map[int]string)
	colors[10] = "red"

	fmt.Println(colors)
}