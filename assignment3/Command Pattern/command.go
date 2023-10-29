package main

import (
	"fmt"
)

type Command interface {
	Execute()
}

type FormulaOneChampionshipLimited struct {}

func (f1cm *FormulaOneChampionshipLimited) Start() {
	fmt.Println("The new season of Formula One Championship starts here at Bakhir International Circut")
}

func (f1cm *FormulaOneChampionshipLimited) Stop() {
	fmt.Println("This is the end of the Formula One Championship season, today will be the last race at Yas-Marina Circut in Abu-Dhabi")
}

func (f1cm *FormulaOneChampionshipLimited) Pause() {
	fmt.Println("The season of Formula One Championship has been paused.")
}


type StartCommand struct {
	championship *FormulaOneChampionshipLimited
}

func (s *StartCommand) Execute() {
	s.championship.Start()
}


type StopCommand struct {
	championship *FormulaOneChampionshipLimited
}

func (s *StopCommand) Execute() {
	s.championship.Stop()
}


type PauseCommand struct {
	championship *FormulaOneChampionshipLimited
}

func (p *PauseCommand) Execute() {
	p.championship.Pause()
}


type ChampionshipControl struct {
	command Command
}

func (c *ChampionshipControl) SetCommand(command Command) {
	c.command = command
}

func (c *ChampionshipControl) ExecuteCommand() {
	c.command.Execute()
}


func main() {
	championship := &FormulaOneChampionshipLimited{}
	start := &StartCommand{championship: championship}
	stop := &StopCommand{championship: championship}
	pause := &PauseCommand{championship: championship}

	championshipControl := ChampionshipControl{}
	
	championshipControl.SetCommand(start)
	championshipControl.ExecuteCommand()

	championshipControl.SetCommand(pause)
	championshipControl.ExecuteCommand()

	championshipControl.SetCommand(stop)
	championshipControl.ExecuteCommand()
}
