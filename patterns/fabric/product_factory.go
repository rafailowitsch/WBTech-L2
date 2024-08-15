package main

type ProductFactory interface {
	CreateProduct() Product
}

type ToyFactory struct{}

func (f *ToyFactory) CreateProduct() Product {
	return &Toy{}
}

type ElectronicsFactory struct{}

func (f *ElectronicsFactory) CreateProduct() Product {
	return &Electronics{}
}
