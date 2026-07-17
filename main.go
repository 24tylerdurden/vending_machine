package main

import (
	"fmt"

	"github.com/pavanIlla/vending_machine/factory"
)

func main() {

	vm := factory.CreateUserVendingMachine()

	vm.SelectItem("Coke")

	vm.InsertCoin(50)

	err := vm.DispenseItem()

	if err != nil {
		fmt.Println("Error:", err)
	}

	vm.SelectItem("Chips")

	vm.DispenseItem()

	vm.SelectItem("Chips")

	vm.DispenseItem()

	vm.SelectItem("Coke")

	fmt.Println(vm.DispenseItem())

}
