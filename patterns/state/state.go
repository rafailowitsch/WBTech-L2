package main

type State interface {
	PressStart(machine *VendingMachine)
	InsertCoin(machine *VendingMachine)
	DispenseItem(machine *VendingMachine)
}
