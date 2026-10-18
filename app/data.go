package main

/*
Declaring data structure
*/

type harvestType string

const (
	Fruit     harvestType = "Fruit"
	Vegetable harvestType = "Vegetable"
	Berry     harvestType = "Berry"
)

type HarvestItem struct {
	name     string
	quantity int
}

type Garden struct {
	itemType map[harvestType][]HarvestItem
}

// For returning items with it's category e.g. in calculations.go "FindSmallestBiggest()"
type CategorizedItem struct {
	category harvestType
	item     HarvestItem
}

/*
End of declaring
*/
