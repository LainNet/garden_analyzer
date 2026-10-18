package main

import "math/rand"

// Iterate through slice of structs for randomize grow of each harvest item
func GrowSimulator(items []HarvestItem) []HarvestItem {
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
