package usecase

import (
	"log"

	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain"
)

// Logger is hook point of internal logs.
var Logger *log.Logger

// AttackByUnterUsecase is written for training of writing clean architecture
type AttackByUnterUsecase interface {
	AttackByHunter(h *domain.Hunter, m *domain.Monster) ([]domain.Material, error)
}

var _ AttackByUnterUsecase = (*Engine)(nil)

// Engine is usecase engine.
type Engine struct {
	domain.HunterRepository
	domain.MonsterRepository
	EventSubscriber
}

// NewEngine returns usecase engine.
func NewEngine(hr domain.HunterRepository, mr domain.MonsterRepository) *Engine {
	return &Engine{
		HunterRepository:  hr,
		MonsterRepository: mr,
	}
}

// AttackByHunter starts to attack by hunter to passed monster.
func (eg *Engine) AttackByHunter(h *domain.Hunter, m *domain.Monster) (profits []domain.Material, err error) {
	logf("[DEBUG]START AttackByHunter: Hunter(%s), Monster(%s)", h.ID, m.ID)
	defer logf("[DEBUG]  END AttackByHunter: Hunter(%s), Monster(%s)", h.ID, m.ID)

	var killed bool
	defer func() {
		if eg.EventSubscriber == nil {
			return
		}
		eg.EventSubscriber.Receive(&EventAttack{
			Killed:  killed,
			Hunter:  h,
			Monster: m,
			Err:     err,
		})
	}()

	profits, killed = h.AttackTo(m)
	if killed {
		h.Materials = append(h.Materials, profits...)
	} else {
		h.DefendFrom(m.Creature)
	}

	if e := eg.HunterRepository.Save(h); e != nil {
		return nil, e
	}
	var e error
	if killed {
		e = eg.MonsterRepository.Remove(m)
	} else {
		eg.MonsterRepository.Save(m)
	}
	if e != nil {
		return nil, e
	}
	return profits, nil
}

func logf(format string, v ...interface{}) {
	if Logger == nil {
		return
	}
	Logger.Printf(format, v...)
}
