package gateway

import (
	"context"

	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain"
)

// Gateway is in usecase, because usecase is domain service.
type Gateway interface {
	HunterRepository() domain.HunterRepository
	MonsterRepository() domain.MonsterRepository
}

// TransactionalGateway is concept in application
type TransactionalGateway interface {
	Gateway

	ContextWithTx(ctx context.Context) (c context.Context, commit, abort context.CancelFunc)
}
