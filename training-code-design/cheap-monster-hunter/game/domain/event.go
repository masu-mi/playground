package domain

type Event interface {
	Name() string
	Error() error
	Summary() string
}

type EventSubscriber interface {
	Receive(e Event)
}
