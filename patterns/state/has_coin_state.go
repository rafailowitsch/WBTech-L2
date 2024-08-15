package main

import "fmt"

type HasCoinState struct{}

func (s *HasCoinState) PressStart(machine *VendingMachine) {
	fmt.Println("Dispensing item...")
	machine.SetState(&NoCoinState{})
}

func (s *HasCoinState) InsertCoin(machine *VendingMachine) {
	fmt.Println("Coin already inserted. Press start to dispense.")
}

func (s *HasCoinState) DispenseItem(machine *VendingMachine) {
	fmt.Println("Press start to dispense the item.")
}
