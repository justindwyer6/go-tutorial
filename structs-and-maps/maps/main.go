package main

import "fmt"

// receiver/type definition
type colorMap map[string]string

func main() {
	// colors := make(map[string]string)
	// var colors map[string]string
	// colors["white"] = "#ffffff"
	// delete(colors, "white")

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"white": "#ffffff",
	// }

	// for color := range colors {
	// 	fmt.Println(color)
	// }

	// receiver/type definition
	var colors colorMap = map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	// printMap(colors)

	// receiver/type definition
	colors.printMap()

	fmt.Println(colors)
}

// func printMap(c map[string]string) {
// 	for color, hex := range c {
// 		fmt.Println("Hex code for", color, "is", hex)
// 	}
// }

// receiver/type definition
func (c colorMap) printMap() {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
