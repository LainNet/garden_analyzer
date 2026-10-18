package main

import (
	"math/rand"
)

// Randomize garden items grows.
func growRandomizer(i *HarvestItem) {
	min := 5
	max := 50
	i.quantity = rand.Intn(max-min+1) + min
}

// start growing items from garden
func (garden Garden) GrowSimulator() {
	for _, items := range garden.itemType {
		for i := range items {
			growRandomizer(&items[i])
		}
	}
}
