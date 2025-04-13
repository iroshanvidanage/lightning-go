package main

import (
	"fmt"
)

func mapsPlay() {
	fmt.Println("Maps")

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}
	fmt.Println(colors)

	printMap(colors)

	colors2 := make(map[int]string)

	colors2[10] = "#ffffff"

	fmt.Println(colors2)
	delete(colors2, 10)
	fmt.Println(colors2)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
