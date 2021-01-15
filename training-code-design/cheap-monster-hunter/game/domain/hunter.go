package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// Hunter is a player or a NPC.
type Hunter struct {
	*Creature

	ID        uuid.UUID  `json:"huter_id"`
	Name      string     `json:"name"`
	Materials []Material `json:"materials"`
}

var humanBase = Creature{
	Life:         100,
	AttackPower:  100,
	DefencePower: 100,
}

var emptyItems = []Material{}

func (h *Hunter) AttackTo(m *Monster) (profits []Material, killed bool) {
	_, killed = h.Creature.AttackTo(m.Creature)
	if !killed {
		return emptyItems, killed
	}
	profits = m.Materials
	if profits == nil {
		profits = emptyItems
	}
	return profits, killed
}

// NewStandardHunter returns a hunter.
func NewStandardHunter(name string) (*Hunter, error) {
	id, e := uuid.NewRandom()
	if e != nil {
		return nil, fmt.Errorf("fail to generate hunter id: %w", e)
	}
	c := humanBase
	return &Hunter{
		ID:       id,
		Name:     name,
		Creature: &c,
	}, nil
}
