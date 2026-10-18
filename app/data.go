package main

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
