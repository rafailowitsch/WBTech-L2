package main

func FactoryUsage() {
	// Создаем фабрику игрушек
	var toyFactory ProductFactory = &ToyFactory{}
	toy := toyFactory.CreateProduct()
	toy.Assemble()
	toy.Package()
	toy.Ship()

	// Создаем фабрику электроники
	var electronicsFactory ProductFactory = &ElectronicsFactory{}
	electronics := electronicsFactory.CreateProduct()
	electronics.Assemble()
	electronics.Package()
	electronics.Ship()
}

func main() {
	FactoryUsage()
}
