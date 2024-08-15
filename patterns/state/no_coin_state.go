package main

import "fmt"

type NoCoinState struct{}

func (s *NoCoinState) PressStart(machine *VendingMachine) {
	fmt.Println("Please insert a coin first.")
}

func (s *NoCoinState) InsertCoin(machine *VendingMachine) {
	fmt.Println("Coin inserted.")
	machine.SetState(&HasCoinState{})
}

func (s *NoCoinState) DispenseItem(machine *VendingMachine) {
	fmt.Println("No coin inserted. Cannot dispense item.")
}
