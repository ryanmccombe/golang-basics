package main

import "fmt"

func main() {
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
	}

	fmt.Println(colors)
	iterateMap(colors)
	a()
	b()
}

func iterateMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// make() maps; add and delete() values
func a() {
	colors := make(map[string]string)

	colors["white"] = "#ffffff"

	fmt.Println(colors)

	delete(colors, "white")

	fmt.Println(colors)
}

// Maps are references
func b() {
	// Just a pointer - not an actual map
	var colors map[string]string

	// Will error - map is null
	// colors["black"] = "#000000"

	initialisedMap := make(map[string]string)

	colors = initialisedMap

	colors["black"] = "#000000"

	fmt.Println(colors)

	// same as above - colors and initialisedMap point to the same map
	fmt.Println(initialisedMap)
}