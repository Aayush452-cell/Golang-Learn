package main

import "fmt"

type Engine struct {
	Type string
}

type Car struct {
	Engine
	Brand string
}

func main() {
	car := Car{
		Engine: Engine{Type: "V8"},
		Brand:  "Audi",
	}
	fmt.Println("Car Brand: ", car.Brand)
	fmt.Println("Engine Type: ", car.Type)
}
