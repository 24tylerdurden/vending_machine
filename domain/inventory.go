package domain

import (
	"fmt"
)

type Inventory struct {
	stock map[string]int
}

func NewInventory() *Inventory {
	return &Inventory{
		stock: make(map[string]int),
	}
}

func (i *Inventory) AddItem(item string, qty int) {
	i.stock[item] += qty
}

func (i *Inventory) IsAvailable(item string) bool {
	qty, exists := i.stock[item]

	return exists && qty > 0
}

func (i *Inventory) ReduceItem(item string) error {
	if !i.IsAvailable(item) {
		return fmt.Errorf(" Item %s is unavailable ", item)
	}

	i.stock[item] -= 1
	return nil
}

func (i *Inventory) GetStock(item string) int {
	return i.stock[item]
}

func (i *Inventory) GetAllStockData() map[string]int {
	// create a temporary map and return it
	copy := make(map[string]int)

	for item, count := range i.stock {
		copy[item] = count
	}

	return copy
}

// conatins Items quantity
