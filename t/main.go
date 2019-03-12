package main

import (
	"fmt"
)

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	Stop string
}

func (this *Car) Start() {
	fmt.Println("Car->Start()")
}

func main() {
	fmt.Println("Start Main func()")
	car := new(Car)
	car.Start()
	// car.Stop()
	car.Stop = "123"
	fmt.Println(car.Stop)
}
