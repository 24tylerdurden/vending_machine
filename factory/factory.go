package factory

import (
	"github.com/pavanIlla/vending_machine/controllers"
	"github.com/pavanIlla/vending_machine/core"
	"github.com/pavanIlla/vending_machine/domain"
	"github.com/pavanIlla/vending_machine/payment"
)

func CreateUserVendingMachine() *controllers.UserVendingMachine {
	inventory := createInventory()
	catalog := createCatalog()
	payment := payment.NewCoinPayment()

	vendingMachine := core.NewVendingMachine(inventory, catalog)

	vendingMachine.SetStrategy(payment)

	userVendingMachine := controllers.NewUserVendingMachine(vendingMachine)

	return userVendingMachine
}

func createInventory() *domain.Inventory {
	inventory := domain.NewInventory()
	inventory.AddItem("Coke", 5)
	inventory.AddItem("Pepsi", 5)
	inventory.AddItem("Chips", 3)
	inventory.AddItem("Water", 10)
	return inventory
}

func createCatalog() *domain.ItemCatalog {
	catalog := domain.NewItemCatalog()
	catalog.AddPrice("Coke", 10)
	catalog.AddPrice("Pepsi", 10)
	catalog.AddPrice("Chips", 20)
	catalog.AddPrice("Water", 5)
	return catalog
}
