package main

import "fmt"

func main() {
	colors := make(map[string]string)

	colors = map[string]string{
		"red":   "#ff00000",
		"green": "#008000",
	}

	colors["white"] = "#ffffff"

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Println("Hex code for", key, "is", val)
	}
}
