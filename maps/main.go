package main

import "fmt"

func main() {
	// Literal declaration
	weeks := map[string]string{
		"sunday":    "Sunday",
		"monday":    "Monday",
		"tuesday":   "Tuesday",
		"wednesday": "Wednesday",
		"thursday":  "Thursday",
		"friday":    "Friday",
		"saturday":  "Saturday",
	}

	fmt.Println(weeks)

	// Declaring map using var
	var colors map[string]int

	colors = make(map[string]int)
	colors["red"] = 1
	colors["green"] = 2
	colors["blue"] = 3

	printMap(colors)

	// Declaring map using make function
	months := make(map[string]int)

	months["Jan"] = 1
	months["Feb"] = 2
	months["Mar"] = 3

	fmt.Println(months)

	// Deleting a key from map
	delete(months, "Mar")
	fmt.Println(months)
}

func printMap(m map[string]int) {
	for k, v := range m {
		fmt.Println("The key is", k, "and the value is", v)
	}
}
