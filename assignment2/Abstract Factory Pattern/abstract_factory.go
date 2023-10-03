package main

import "fmt"

// Interface for all motorcycles
type IMotorcycle interface {
	GetType() string
}

// Interface for factory
type IMotorcycleFactory interface {
	CreateMotorcycle() IMotorcycle
}

// Sportbike model
type SportMotorcycle struct{}

func (sportMotorcycle *SportMotorcycle) GetType() string {
	return "SportBike yahoooooooo"
}

// Cruiser model
type CruiserMotorcycle struct{}

func (cruiserMotorcycle *CruiserMotorcycle) GetType() string {
	return "Cruiser (chill out after sport xd)"
}


// Sportbike factory
type SportMotorcycleFactory struct{}

func (sportMotorcycleFactory *SportMotorcycleFactory) CreateMotorcycle() IMotorcycle {
    return &SportMotorcycle{}
}

// Cruiser factory
type CruiserMotorcycleFactory struct{}

func (cruiserMotorcycleFactory *CruiserMotorcycleFactory) CreateMotorcycle() IMotorcycle {
    return &CruiserMotorcycle{}
}



func main() {
	sportFactory := SportMotorcycleFactory{}
	kawasakiNinja := sportFactory.CreateMotorcycle()
	fmt.Println(kawasakiNinja.GetType())

	cruiserFactory := CruiserMotorcycleFactory{}
	harleyDavidson := cruiserFactory.CreateMotorcycle()
	fmt.Println(harleyDavidson.GetType())
}
