package mediator

// Train is interface
type Train interface {
	RequestArrival()
	Depart()
	PermitArrival()
}
