package domain

import (
	"context"

	"github.com/google/uuid"
)

// HunterRepository persistency manager of Hunter
type HunterRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Hunter, error)
	Save(ctx context.Context, hunter *Hunter) error
	Remove(ctx context.Context, hunter *Hunter) error
}
