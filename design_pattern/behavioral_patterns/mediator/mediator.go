package mediator

/*
	Mediator is a behavioral design pattern that lets you reduce chaotic dependencies between objects.
	The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object.
	The Mediator interface declares methods of communication with components, which usually include just a single notification method.

collaborate: cong tac
*/
type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}
