package main

import "fmt"

func main() {
	garden := Garden{
		itemType: make(map[harvestType][]HarvestItem),
	}

	garden.itemType[Fruit] = []HarvestItem{
		{"Apple", 0},
		{"Banana", 0},
		{"Orange", 0},
		{"Mango", 0},
		{"Kiwi", 0},
	}
	garden.itemType[Vegetable] = []HarvestItem{
		{"Cucumber", 0},
		{"Beets", 0},
		{"Spinach", 0},
		{"Potato", 0},
		{"Pumkin", 0},
	}
	garden.itemType[Berry] = []HarvestItem{
		{"Strawberry", 0},
		{"Redcurrants", 0},
		{"Blueberries", 0},
		{"Blackcurrants", 0},
		{"Elderberries", 0},
	}

	garden.GrowSimulator()
	smallest, biggest := garden.FindSmallestBiggest()
	fmt.Printf("The smallest harvestable: '%s' — '%s' with '%d' items\n", smallest.category, smallest.item.name, smallest.item.quantity)
	fmt.Printf("The biggest harvestable: '%s' — '%s' with '%d' items\n\n", biggest.category, biggest.item.name, biggest.item.quantity)

	fmt.Printf("Total harvestables in garden: '%d'\n", garden.TotalGardenItems())
	PrintSortedTotals(garden)

}
