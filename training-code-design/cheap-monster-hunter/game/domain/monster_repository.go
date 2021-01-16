package domain

import (
	"context"

	"github.com/google/uuid"
)

// MonsterRepository persistency manager of Monster
type MonsterRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Monster, error)
	Save(ctx context.Context, hunter *Monster) error
	Remove(ctx context.Context, hunter *Monster) error
}
