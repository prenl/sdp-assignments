package main

import "fmt"

// Formula driver and team models
type FormulaDriver struct {
	name string
	number int
}

type FormulaTeam struct {
	name string
	firstDriver FormulaDriver
	secondDriver FormulaDriver
}

// Method to 'notify' the drivers of the team
func (team FormulaTeam) Update(msg string) {
	fmt.Printf("Message for %s and %s:\n%s\n\n", team.firstDriver.name, team.secondDriver.name, msg)
}

// Observer and race controller
type IObserver interface {
	Update(string)
}

type RaceControl struct {
	observers []IObserver
}

// Attach method for adding observers to race controller
func (raceControl *RaceControl) Attach(observer IObserver) {
	raceControl.observers = append(raceControl.observers, observer)
}

// Notify for a yellow flag (Danger on the track
func (raceControl *RaceControl) NotifyYellowFlag() {
	for _, observer := range raceControl.observers {
		observer.Update("Yellow flag! Slow down")
	}
}

// Notify for a red flag (Race stop due to weather or inchident)
func (raceControl *RaceControl) NotifyRedFlag() {
	for _, observer := range raceControl.observers {
		observer.Update("Red flag! Red flag! Slow down... Get in the pit lane")
	}
}

// Notify for a Safety Car Mode (Car that slows race's pace due to weather or inchident)
func (raceControl *RaceControl) NotifySCMode() {
	for _, observer := range raceControl.observers {
		observer.Update("Safety Car has deployed, pick up the pace to catch it")
	}
}



func main() {
	redBullRacing := FormulaTeam{ 
		name: "Oracle Red Bull Racing Formula One® Team",
		firstDriver: FormulaDriver{name: "Max Verstappen", number: 1},
		secondDriver: FormulaDriver{name: "Checo Perez", number: 11},
	}

	astonMartin := FormulaTeam{
		name: "Aston Martin Aramco Cognizant Formula One® Team",
		firstDriver: FormulaDriver{name: "Fernando Alonso", number: 14},
		secondDriver: FormulaDriver{name: "Lance Stroll", number: 18},
	}

	mcLaren := FormulaTeam{
		name: "McLaren Formula One® Team",
		firstDriver: FormulaDriver{name: "Lando Norris", number: 4},
		secondDriver: FormulaDriver{name: "Oscar Piastri", number: 81},
	}

	raceControl := RaceControl{}

	raceControl.Attach(redBullRacing)
	raceControl.Attach(astonMartin)
	raceControl.Attach(mcLaren)

	raceControl.NotifySCMode()
}
