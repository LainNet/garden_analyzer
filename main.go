package main

import (
	"fmt"
	"math/rand"
)

// Declaring data structure
type HarvestCategory string

const (
	Fruit     HarvestCategory = "Fruit"
	Vegetable HarvestCategory = "Vegetable"
	Berry     HarvestCategory = "Berry"
)

type HarvestItem struct {
	name     string
	quantity int
	category HarvestCategory
}

// End of declaring

func main() {

	Garden := []HarvestItem{
		{"Apple", 0, Fruit},
		{"Banana", 0, Fruit},
		{"Orange", 0, Fruit},
		{"Mango", 0, Fruit},
		{"Kiwi", 0, Fruit},

		{"Cucumber", 0, Vegetable},
		{"Beets", 0, Vegetable},
		{"Spinach", 0, Vegetable},
		{"Potato", 0, Vegetable},
		{"Pumkin", 0, Vegetable},

		{"Strawberry", 0, Berry},
		{"Redcurrants", 0, Berry},
		{"Blueberries", 0, Berry},
		{"Blackcurrants", 0, Berry},
		{"Elderberries", 0, Berry},
	}

	growned := growSimulator(Garden)
	smallestBiggest := findSmallestBiggest(growned)

	// Result:
	fmt.Printf("Smallest produced: '%s' with quantity of '%d' from '%s' garden", smallestBiggest[0].name, smallestBiggest[0].quantity, smallestBiggest[0].category)
	fmt.Println()
	fmt.Printf("Biggest produced: '%s' with quantity of '%d' from '%s' garden", smallestBiggest[1].name, smallestBiggest[1].quantity, smallestBiggest[1].category)
}

// Find and return smallest & biggest harvest item from garden
func findSmallestBiggest(items []HarvestItem) []HarvestItem {
	minId, maxId := 0, 0

	for i, item := range items {
		if item.quantity < items[minId].quantity {
			minId = i
		}
		if item.quantity > items[maxId].quantity {
			maxId = i
		}

	}

	return []HarvestItem{items[minId], items[maxId]}
}

// Iterate through slice of structs for randomize grow of each harvest item
func growSimulator(items []HarvestItem) []HarvestItem {
	for i := range len(items) {
		growRandomizer(&items[i])
	}
	return items
}

// Randomize garden items grows.
func growRandomizer(i *HarvestItem) {
	min := 5
	max := 50
	i.quantity = rand.Intn(max-min+1) + min
}
