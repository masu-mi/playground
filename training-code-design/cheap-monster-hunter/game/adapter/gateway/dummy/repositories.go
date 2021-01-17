package dummy

import (
	"context"

	"github.com/google/uuid"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain"
)

type HunterRepo struct {
	repo map[uuid.UUID]*domain.Hunter
}

func NewHunterRepo() *HunterRepo {
	return &HunterRepo{repo: map[uuid.UUID]*domain.Hunter{}}
}

var _ domain.HunterRepository = (*HunterRepo)(nil)

func (hr *HunterRepo) FindByID(_ context.Context, id uuid.UUID) (*domain.Hunter, error) {
	if v, ok := hr.repo[id]; ok {
		return v, nil
	}
	return nil, &domain.ErrNotFound{ID: id}
}

func (hr *HunterRepo) Save(_ context.Context, h *domain.Hunter) error {
	hr.repo[h.ID] = h
	return nil
}
func (hr *HunterRepo) Remove(_ context.Context, h *domain.Hunter) error {
	delete(hr.repo, h.ID)
	return nil
}

type MonsterRepo struct{}

var _ domain.MonsterRepository = (*MonsterRepo)(nil)

func (hr *MonsterRepo) FindByID(_ context.Context, id uuid.UUID) (*domain.Monster, error) {
	u, _ := uuid.NewUUID()
	return &domain.Monster{
		Name:     "mock monster",
		ID:       u,
		Creature: &domain.Creature{Life: 5, AttackPower: 100, DefencePower: 1},
		Materials: []domain.Material{
			{Name: "酒", Rarity: 100},
			{Name: "米", Rarity: 1000},
		},
	}, nil
}

func (hr *MonsterRepo) Save(_ context.Context, _ *domain.Monster) error {
	return nil
}
func (hr *MonsterRepo) Remove(_ context.Context, _ *domain.Monster) error {
	return nil
}
