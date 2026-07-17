package controllers

import (
	"github.com/pavanIlla/vending_machine/core"
)

type UserVendingMachine struct {
	machine *core.VendingMachine
}

func NewUserVendingMachine(machine *core.VendingMachine) *UserVendingMachine {
	return &UserVendingMachine{
		machine: machine,
	}
}

func (vm *UserVendingMachine) SelectItem(item string) error {
	return vm.machine.SelectItem(item)
}

func (vm *UserVendingMachine) InsertCoin(amount int) error {
	return vm.machine.InsertCoin(amount)
}

func (vm *UserVendingMachine) DispenseItem() error {
	return vm.machine.DispenseItem()
}
