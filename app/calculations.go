package main

import (
	"fmt"
	"sort"
)

func (garden Garden) FindSmallestBiggest() (CategorizedItem, CategorizedItem) {
	var minResult, maxResult CategorizedItem
	initialzied := false
	for itemType, items := range garden.itemType {
		for _, item := range items {
			if !initialzied {
				minResult = CategorizedItem{category: itemType, item: item}
				maxResult = CategorizedItem{category: itemType, item: item}
				initialzied = true
				continue
			}

			if item.quantity < minResult.item.quantity {
				minResult = CategorizedItem{category: itemType, item: item}
			}
			if item.quantity > maxResult.item.quantity {
				maxResult = CategorizedItem{category: itemType, item: item}
			}
		}
	}
	return minResult, maxResult
}

// Calculating total of garden harvestables
func (garden Garden) TotalGardenItems() int {
	total := 0

	for _, items := range garden.itemType {
		for _, item := range items {
			total += item.quantity
		}
	}
	return total
}

// Calculating total by category
func (garden Garden) totalPerCategory() map[harvestType]int {
	totals := make(map[harvestType]int)

	for itemType, items := range garden.itemType {
		sum := 0
		for _, item := range items {
			sum += item.quantity
		}
		totals[itemType] = sum
	}

	return totals
}

// Printing Total of harvestables by category

func PrintSortedTotals(garden Garden) {
	total := garden.totalPerCategory()

	// Collect the keys (harvest types) into a slice
	var keys []harvestType
	for itemType := range total {
		keys = append(keys, itemType)
	}

	// Sort the keys alphabetically
	sort.Slice(keys, func(i, j int) bool {
		return string(keys[i]) < string(keys[j])
	})

	// Print them in order
	fmt.Println("Total harvestables by category: ")
	for _, itemType := range keys {
		fmt.Printf("\t%s: %d items\n", itemType, total[itemType])
	}
}
