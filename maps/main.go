package main

import "fmt"

func main() {
	colors := map[string]string{
		"black":  "000000",
		"red":    "ff0000",
		"green":  "00ff00",
		"blue":   "0000ff",
		"OOPS!!": "Wrong Entry!",
	}

	delete(colors, "OOPS!!")

	printMap(colors)

	colors["white"] = "ffffff"

	fmt.Println(colors["white"], " is the hex of white")
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex of ", color, " is #", hex)
	}
}
