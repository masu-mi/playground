package dummy

import (
	"github.com/google/uuid"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain"
)

type HunterRepo struct{}

var _ domain.HunterRepository = (*HunterRepo)(nil)

func (hr *HunterRepo) FindByID(id uuid.UUID) (*domain.Hunter, error) {
	u, _ := uuid.NewUUID()
	return &domain.Hunter{
		Name:     "mock",
		ID:       u,
		Creature: &domain.Creature{Life: 100, AttackPower: 10, DefencePower: 20},
	}, nil
}

func (hr *HunterRepo) Save(hunterRepo *domain.Hunter) error {
	return nil
}
func (hr *HunterRepo) Remove(hunterRepo *domain.Hunter) error {
	return nil
}

type MonsterRepo struct{}

var _ domain.MonsterRepository = (*MonsterRepo)(nil)

func (hr *MonsterRepo) FindByID(id uuid.UUID) (*domain.Monster, error) {
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

func (hr *MonsterRepo) Save(_ *domain.Monster) error {
	return nil
}
func (hr *MonsterRepo) Remove(_ *domain.Monster) error {
	return nil
}
