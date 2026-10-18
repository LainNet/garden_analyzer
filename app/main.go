package main

import (
	"fmt"
)

func main() {

	garden := []HarvestItem{
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

	growned := GrowSimulator(garden)
	smallestBiggest := FindSmallestBiggest(growned)

	// Result:
	fmt.Printf("Smallest produced: '%s' with quantity of '%d' from '%s' garden", smallestBiggest[0].name, smallestBiggest[0].quantity, smallestBiggest[0].category)
	fmt.Println()
	fmt.Printf("Biggest produced: '%s' with quantity of '%d' from '%s' garden", smallestBiggest[1].name, smallestBiggest[1].quantity, smallestBiggest[1].category)
	fmt.Println()
	fmt.Printf("Total Harvestable items in garden: %d", TotalGardenItems(growned))
}

// Declaring data structure
type harvestCategory string

const (
	Fruit     harvestCategory = "Fruit"
	Vegetable harvestCategory = "Vegetable"
	Berry     harvestCategory = "Berry"
)

type HarvestItem struct {
	name     string
	quantity int
	category harvestCategory
}

// End of declaring
