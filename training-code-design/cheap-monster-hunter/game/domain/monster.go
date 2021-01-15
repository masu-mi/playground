package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// Monster is a NPC is player's target.
type Monster struct {
	*Creature

	ID        uuid.UUID  `json:"huter_id"`
	Name      string     `json:"name"`
	Materials []Material `json:"hunted_materials"`
}

// NewMonster returns a monster.
func NewMonster(name string) (*Monster, error) {
	id, e := uuid.NewRandom()
	if e != nil {
		return nil, fmt.Errorf("fail to generate Monster id: %w", e)
	}
	return &Monster{ID: id, Name: name}, nil
}

// MonsterFactory returns monster of race.s
func MonsterFactory(name string, level int, race CreatureRace) (*Monster, error) {
	m, e := NewMonster(name)
	if e != nil {
		return nil, fmt.Errorf("fail to create monster with %w", e)
	}
	m.Creature = race.NewCreature(level)
	return m, nil
}
