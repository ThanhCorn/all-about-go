package decorator

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

type IBread interface {
	Cost() float64
}

type Bread struct {
}

func (b *Bread) Cost() float64 {
	return 5.0
}

type PorkRoll struct {
	Bread IBread
}

type Omelet struct {
	Bread IBread
}

func (p *PorkRoll) Cost() float64 {
	cost := 5.0
	fmt.Printf("Add some pork roll to the bread price %f\n", cost)
	return p.Bread.Cost() + cost
}

func (p *Omelet) Cost() float64 {
	cost := 10.0
	fmt.Printf("Add some omelet to the bread price %f\n", cost)
	return p.Bread.Cost() + cost
}

func TestOrderBread(t *testing.T) {
	bread := &Bread{}
	porkRoll := &PorkRoll{
		Bread: bread,
	}
	omelet := &Omelet{
		Bread: porkRoll,
	}
	omeletPorkRollBread := omelet.Cost()
	assert.Equal(t, omeletPorkRollBread, 20.0)

}
