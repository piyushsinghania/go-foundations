package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"Red":    "#FF0000",
		"Green":  "#00FF00",
		"Blue":   "#0000FF",
		"Yellow": "#FFFF00",
		"Purple": "#800080",
	}

	colors["White"] = "#FFFFFF"
	fmt.Println(colors)

	delete(colors, "Red")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%s: %s\n", color, hex)
	}
}
