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

	fruits := [5]HarvestItem{
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

}

func growRandomizer(h *HarvestItem) {
	min := 5
	max := 50
	h.quantity = rand.Intn(max-min+1) + min
}
