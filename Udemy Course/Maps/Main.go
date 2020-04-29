package main

import "fmt"

func main() {
	//maps are like objects but differ from structs
	// keys and values are statically typed (ie they must be the same type)
	// there are multiple ways to declare a map in go

	// bog standard syntax all the keys are strings and all the values are strings also
	// colors := map[string]string{
	// 	"red":  "ff0000",
	// 	"blue": "0000FF",
	// }

	// var syntax for maps better than first map as easier to add and remove later
	// var colors map[string]string

	// third way of creating map better than first map as easier to add and remove later

	colors := make(map[string] /* keys*/ string /* values */)

	// here we add an entry to the map
	colors["white"] = "#fffff"
	colors["Green"] = "#fff23"
	colors["Red"] = "#ff0000"

	// here we delete the white entry using its key
	delete(colors, "white")

	//to access individual key we use the [] syntax

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is ", hex)
	}
}
