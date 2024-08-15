package main

import "fmt"

type Product interface {
	Assemble()
	Package()
	Ship()
}

type Toy struct{}

func (t *Toy) Assemble() {
	fmt.Println("Assembling toy")
}

func (t *Toy) Package() {
	fmt.Println("Packaging toy")
}

func (t *Toy) Ship() {
	fmt.Println("Shipping toy")
}

type Electronics struct{}

func (e *Electronics) Assemble() {
	fmt.Println("Assembling electronics")
}

func (e *Electronics) Package() {
	fmt.Println("Packaging electronics")
}

func (e *Electronics) Ship() {
	fmt.Println("Shipping electronics")
}
