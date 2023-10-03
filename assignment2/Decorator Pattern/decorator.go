package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // Seed the random generator
}

// Default Car Pitstop Interface
type ICarPitstop interface {
	Refuel() string
	GetTime() float64
}

// Car Pitstop
type BasicCarPitstop struct{}

func (basicCarPitstop *BasicCarPitstop) Refuel() string {
	return "Car refueled!"
}

func (basicCarPitstop *BasicCarPitstop) GetTime() float64 {
	return float64(rand.Intn(60)) // Random number 0 - 60
}

// Pit stop Tyre Change
type TireChange struct {
	carPitstop ICarPitstop
}

func (tireChange *TireChange) Refuel() string {
	return tireChange.carPitstop.Refuel() + " Tyres changed!"
}

func (tireChange *TireChange) GetTime() float64 {
	return tireChange.carPitstop.GetTime() + float64(rand.Intn(12))
}

// Pit stop Driver Change
type DriverChange struct {
	carPitstop ICarPitstop
}

func (driverChange *DriverChange) Refuel() string {
	return driverChange.carPitstop.Refuel() + " Driver changed!"
}

func (driverChange *DriverChange) GetTime() float64 {
	return driverChange.carPitstop.GetTime() + 45.0
}

func main() {
	var service ICarPitstop
	service = &BasicCarPitstop{}
	service = &TireChange{carPitstop: service}
	service = &DriverChange{carPitstop: service}

	fmt.Println("\n24 heures Du Mans Pit Stop!")
	fmt.Println(service.Refuel())
	fmt.Printf("%d seconds wasted\n\n", int64(service.GetTime()))
}
