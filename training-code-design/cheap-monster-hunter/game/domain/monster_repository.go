package domain

import "github.com/google/uuid"

// MonsterRepository persistency manager of Monster
type MonsterRepository interface {
	FindByID(id uuid.UUID) (*Monster, error)
	Save(hunter *Monster) error
	Remove(hunter *Monster) error
}
