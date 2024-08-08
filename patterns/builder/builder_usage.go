package main

import "fmt"

func Builderusage() {
	gamingBuilder := &GamingComputerBuilder{}
	officeBuilder := &OfficeComputerBuilder{}

	director := Director{}

	director.SetBuilder(gamingBuilder)
	gamingComputer := director.BuildComputer()
	fmt.Printf("Gaming Computer: %+v\n", gamingComputer)

	director.SetBuilder(officeBuilder)
	officeComputer := director.BuildComputer()
	fmt.Printf("Office Computer: %+v\n", officeComputer)
}
