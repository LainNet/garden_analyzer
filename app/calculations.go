package main

// Find and return smallest & biggest harvest item from garden
func FindSmallestBiggest(items []HarvestItem) []HarvestItem {
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

// Return total of garden harvestable items.
func TotalGardenItems(items []HarvestItem) int {
	total := 0
	for _, item := range items {
		total += item.quantity
	}
	return total
}
