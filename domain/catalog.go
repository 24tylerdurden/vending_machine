package domain

import "fmt"

type ItemCatalog struct {
	catalog map[string]int
}

func NewItemCatalog() *ItemCatalog {
	return &ItemCatalog{
		catalog: make(map[string]int),
	}
}

func (i *ItemCatalog) AddPrice(item string, cost int) error {

	if cost < 0 {
		return fmt.Errorf("Invalid %d price", cost)
	}

	i.catalog[item] = cost
	return nil
}

func (i *ItemCatalog) GetPrice(item string) (int, error) {
	if !i.HasItem(item) {
		return 0, fmt.Errorf("Product doesn't exists")
	}

	return i.catalog[item], nil
}

func (i *ItemCatalog) HasItem(item string) bool {
	_, exists := i.catalog[item]

	return exists
}

// contains Item price
