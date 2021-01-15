package usecase

import (
	"fmt"

	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain"
)

type Event interface {
	Name() string
	Error() error
	Summary() string
}

type EventSubscriber interface {
	Receive(e Event)
}

type EventAttack struct {
	Killed bool

	*domain.Hunter
	*domain.Monster

	Err error
}

func (ea *EventAttack) Name() string { return "EventAttack" }
func (ea *EventAttack) Error() error { return ea.Err }
func (ea *EventAttack) Summary() string {
	return fmt.Sprintf("%s attacks to %s", ea.Hunter.Name, ea.Monster.Name)
}
