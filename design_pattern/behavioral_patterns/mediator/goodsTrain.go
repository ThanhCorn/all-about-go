package mediator

import "fmt"

type GoodsTrain struct {
	Mediator Mediator
}

func (g *GoodsTrain) RequestArrival() {
	if g.Mediator.canArrive(g) {
		fmt.Printf("\nGoodsTrain Arrival")
	} else {
		fmt.Printf("\nGoodsTrain Waiting")
	}
}

func (g *GoodsTrain) Depart() {
	fmt.Printf("\nGoodsTrain Departing...")
	g.Mediator.notifyAboutDeparture()
}

func (g *GoodsTrain) PermitArrival() {
	fmt.Printf("\nGoodsTrain arrival permitted")
}
