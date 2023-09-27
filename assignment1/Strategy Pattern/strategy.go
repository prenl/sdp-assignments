package main

import "fmt"

type IDriveStrategy interface {
	Drive()
}

// Structure for fast and slow driving
type DriveFast struct {}
type DriveSlow struct {}
type DriveNormal struct {}


// Implementing Drive method from IDriveable
func (driveFast DriveFast) Drive() {
	fmt.Println("Driving fast!")
}

func (driveSlow DriveSlow) Drive() {
	fmt.Println("Driving slow")
}

func (driveNormal DriveNormal) Drive() {
	fmt.Println("Driving as usual")
}


// Car Model
type Car struct {
	name string
	kW int
	driveStrategy IDriveStrategy
}

// Method for Car model
func (car Car) Drive() {
	car.driveStrategy.Drive()
}

func (car Car) Info() {
	fmt.Printf("---- Car %s ----\nPower: %d kW\nDrive Strategy: %T\n", car.name, car.kW, car.driveStrategy)
}



func main() {
	myCar := Car{
		name: "Audi RS6 C8",
		kW: 450,
		driveStrategy: DriveFast{},
	}

	myCar.Info()
	myCar.Drive()
	myCar.driveStrategy = DriveSlow{}
	myCar.Info()
	myCar.Drive()

}
