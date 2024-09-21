package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {
	station := NewStationManager()
	passengerTrain := PassengerTrain{
		Mediator: station,
	}
	goodsTrain := GoodsTrain{
		Mediator: station,
	}

	goodsTrain.RequestArrival()
	passengerTrain.RequestArrival()
	goodsTrain.Depart()
	passengerTrain.Depart()
}
