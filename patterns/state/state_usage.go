package main

func StateUsage() {
	machine := &VendingMachine{state: &NoCoinState{}}

	machine.InsertCoin()
	machine.PressStart()
	machine.DispenseItem()

	machine.InsertCoin()
	machine.InsertCoin()
	machine.PressStart()
	machine.DispenseItem()
}

func main() {
	StateUsage()
}
