package dummy

import (
	"context"
	"fmt"

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

func getTx(ctx context.Context) *tx {
	txV, ok := ctx.Value(txKey{}).(*tx)
	if ok {
		fmt.Printf("%v\n", txV)
		return txV
	}
	return &tx{} // default value: like auto commit db
}

func (hr *HunterRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Hunter, error) {
	_ = getTx(ctx)
	if h, ok := hr.repo[id]; ok {
		return h, nil
	}
	return nil, &domain.ErrNotFound{ID: id}
}

func (hr *HunterRepo) Save(ctx context.Context, h *domain.Hunter) error {
	_ = getTx(ctx)
	hr.repo[h.ID] = h
	return nil
}
func (hr *HunterRepo) Remove(ctx context.Context, h *domain.Hunter) error {
	_ = getTx(ctx)
	delete(hr.repo, h.ID)
	return nil
}

type MonsterRepo struct{}

var _ domain.MonsterRepository = (*MonsterRepo)(nil)

func (hr *MonsterRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Monster, error) {
	_ = getTx(ctx)
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

func (hr *MonsterRepo) Save(ctx context.Context, _ *domain.Monster) error {
	_ = getTx(ctx)
	return nil
}
func (hr *MonsterRepo) Remove(ctx context.Context, _ *domain.Monster) error {
	_ = getTx(ctx)
	return nil
}
