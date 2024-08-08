package main

import "fmt"

func main() {
	components := []Component{
		&DB{"localhost:5432", "postgres", "user:password"},
		&Server{"localhost", "5s"},
	}

	infovisitors := &InfoVisitors{}

	for _, visitor := range components {
		visitor.Accept(infovisitors)
		fmt.Println(infovisitors.info)
	}
}
