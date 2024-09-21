package mediator

import "fmt"

type PassengerTrain struct {
	Mediator Mediator
}

func (p *PassengerTrain) RequestArrival() {
	if p.Mediator.canArrive(p) {
		fmt.Printf("\nPassengerTrain Arrival")
	} else {
		fmt.Printf("\nPassengerTrain Waiting")
	}
}

func (p *PassengerTrain) Depart() {
	fmt.Printf("\nPassengerTrain Departing...")
	p.Mediator.notifyAboutDeparture()
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Printf("\nPassengerTrain arrival permitted")
}
