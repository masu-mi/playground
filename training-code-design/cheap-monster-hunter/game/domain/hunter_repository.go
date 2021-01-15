package domain

import (
	"github.com/google/uuid"
)

// HunterRepository persistency manager of Hunter
type HunterRepository interface {
	FindByID(id uuid.UUID) (*Hunter, error)
	Save(hunter *Hunter) error
	Remove(hunter *Hunter) error
}
