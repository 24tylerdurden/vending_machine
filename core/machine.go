package core

import (
	"fmt"

	"github.com/pavanIlla/vending_machine/domain"
	"github.com/pavanIlla/vending_machine/interfaces"
)

type VendingMachine struct {
	inventory    *domain.Inventory
	catalog      *domain.ItemCatalog
	paymentStrat interfaces.PaymentStrategy
	selectedItem string
}

func NewVendingMachine(inventory *domain.Inventory, catalog *domain.ItemCatalog) *VendingMachine {
	return &VendingMachine{
		inventory: inventory,
		catalog:   catalog,
	}
}

func (vm *VendingMachine) SetStrategy(start interfaces.PaymentStrategy) {
	vm.paymentStrat = start
}

func (vm *VendingMachine) SelectItem(item string) error {
	if !vm.catalog.HasItem(item) {
		return fmt.Errorf("item %s doesn't exists", item)
	}

	vm.selectedItem = item
	price, _ := vm.catalog.GetPrice(item)
	fmt.Printf("Selected: %s, Price: %d\n", item, price)
	return nil
}

func (vm *VendingMachine) InsertCoin(amount int) error {
	if vm.paymentStrat == nil {
		return fmt.Errorf("no payment method set")
	}
	vm.paymentStrat.AddMoney(amount)
	return nil
}

func (vm *VendingMachine) CancelTransaction() {

}

func (vm *VendingMachine) DispenseItem() error {
	if vm.selectedItem == "" {
		return fmt.Errorf("no item selected")
	}

	if vm.paymentStrat == nil {
		return fmt.Errorf("no payment method has been set")
	}

	if !vm.inventory.IsAvailable(vm.selectedItem) {
		return fmt.Errorf("out of stock: %s", vm.selectedItem)
	}

	price, err := vm.catalog.GetPrice(vm.selectedItem)

	if err != nil {
		return err
	}

	success, err := vm.paymentStrat.Pay(price)

	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("insufficient money. Need: %d", price)
	}

	if err := vm.inventory.ReduceItem(vm.selectedItem); err != nil {
		return err
	}

	if change := vm.paymentStrat.GetBalance(); change > 0 {
		fmt.Printf("Change returned: %d\n", change)
	}

	vm.selectedItem = ""

	return nil

}
