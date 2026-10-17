package main

import (
	"fmt"
	"math/rand"
)

type HarvestItem struct {
	name     string
	quantity int
}

func main() {

	fruits := []HarvestItem{
		{"apple", 0},
		{"banana", 0},
		{"orange", 0},
		{"mango", 0},
		{"Kiwi", 0},
	}

	// growRandomizer(&fruits[2])
	// fmt.Println(fruits[2])

	for i := range len(fruits) {
		growRandomizer(&fruits[i])
	}

	fmt.Println(fruits)
	fmt.Println(findBiggestSmallest(fruits))
}

// Finding and returning the smallest and biggest growned product in garden
func findBiggestSmallest(items []HarvestItem) []HarvestItem {
	smallest, biggest := items[0].quantity, items[0].quantity
	var bigName, smallName string
	for _, i := range items {
		if i.quantity > biggest {
			biggest = i.quantity
			bigName = i.name
		}
		if i.quantity < smallest {
			smallest = i.quantity
			smallName = i.name
		}
	}
	return []HarvestItem{
		{bigName, biggest},
		{smallName, smallest},
	}
}

// Randomize garden items grows.
func growRandomizer(h *HarvestItem) {
	min := 5
	max := 50
	h.quantity = rand.Intn(max-min+1) + min
}
