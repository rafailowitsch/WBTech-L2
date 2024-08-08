package main

type ComputerBuilder interface {
	SetCPU() ComputerBuilder
	SetRAM() ComputerBuilder
	SetStorage() ComputerBuilder
	SetGPU() ComputerBuilder
	GetComputer() Computer
}

type GamingComputerBuilder struct {
	computer Computer
}

func (b *GamingComputerBuilder) SetCPU() ComputerBuilder {
	b.computer.CPU = "Intel i9"
	return b
}

func (b *GamingComputerBuilder) SetRAM() ComputerBuilder {
	b.computer.RAM = "32GB"
	return b
}

func (b *GamingComputerBuilder) SetStorage() ComputerBuilder {
	b.computer.Storage = "1TB SSD"
	return b
}

func (b *GamingComputerBuilder) SetGPU() ComputerBuilder {
	b.computer.GPU = "NVIDIA RTX 3080"
	return b
}

func (b *GamingComputerBuilder) GetComputer() Computer {
	return b.computer
}

type OfficeComputerBuilder struct {
	computer Computer
}

func (b *OfficeComputerBuilder) SetCPU() ComputerBuilder {
	b.computer.CPU = "Intel i5"
	return b
}

func (b *OfficeComputerBuilder) SetRAM() ComputerBuilder {
	b.computer.RAM = "16GB"
	return b
}

func (b *OfficeComputerBuilder) SetStorage() ComputerBuilder {
	b.computer.Storage = "512GB SSD"
	return b
}

func (b *OfficeComputerBuilder) SetGPU() ComputerBuilder {
	b.computer.GPU = "Integrated"
	return b
}

func (b *OfficeComputerBuilder) GetComputer() Computer {
	return b.computer
}
