package main

type VendingMachine struct {
	state State
}

func (m *VendingMachine) SetState(state State) {
	m.state = state
}

func (m *VendingMachine) PressStart() {
	m.state.PressStart(m)
}

func (m *VendingMachine) InsertCoin() {
	m.state.InsertCoin(m)
}

func (m *VendingMachine) DispenseItem() {
	m.state.DispenseItem(m)
}
