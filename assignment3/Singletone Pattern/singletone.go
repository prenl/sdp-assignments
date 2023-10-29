package main

import (
	"fmt"
	"sync"
)

type SafetyCar struct {
	PilotName string
	Status int
}

var safetyCar *SafetyCar
var safetyCarOnce sync.Once

func GetSafetyCar() *SafetyCar {
	safetyCarOnce.Do(func() {
		safetyCar = &SafetyCar{}
	})
	return safetyCar
}

func main() {
	safetyCar1 := GetSafetyCar()
	safetyCar1.PilotName = "Bernd Mayländer"

	safetyCar2 := GetSafetyCar()

	fmt.Println("№1 SafetyCar Pilot: " + safetyCar1.PilotName)
	fmt.Println("№2 SafetyCar Pilot: " + safetyCar2.PilotName)
	
	if (safetyCar1 == safetyCar2) {
		fmt.Println("The Safety Cars have the same instance!")
	} else {
		fmt.Println("The Safety Cars don't have the same instance!")
	}
}
