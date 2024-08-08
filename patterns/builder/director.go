package main

type Director struct {
	builder ComputerBuilder
}

func (d *Director) SetBuilder(b ComputerBuilder) {
	d.builder = b
}

func (d *Director) BuildComputer() Computer {
	return d.builder.SetCPU().
		SetRAM().
		SetStorage().
		SetGPU().
		GetComputer()
}
