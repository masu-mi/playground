package dummy

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/adapter/gateway"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain"
	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/usecase"
)

// Gateway is Engine factory.
type Gateway struct {
	*HunterRepo
	*MonsterRepo

	EventLogger usecase.EventSubscriber
}

var _ gateway.TransactionalGateway = (*Gateway)(nil)

func (g *Gateway) HunterRepository() domain.HunterRepository {
	return g.HunterRepo
}

func (g *Gateway) MonsterRepository() domain.MonsterRepository {
	return g.MonsterRepo
}

func (g *Gateway) ContextWithTx(p context.Context) (ch context.Context, commit, abort context.CancelFunc) {
	txV := &tx{}
	ch = context.WithValue(p, txKey{}, txV)
	var cancel context.CancelFunc
	ch, cancel = context.WithCancel(ch)
	go func() {
		select {
		case <-ch.Done():
		}
	}()
	return ch, func() { txV.Commit(); cancel() }, func() { txV.Abort(); cancel() }
}

// Engine returns engine.
func (g *Gateway) Engine(c context.Context) (*usecase.Engine, context.Context) {
	eng := usecase.NewEngine(&HunterRepo{}, &MonsterRepo{})
	eng.EventSubscriber = usecase.NewEventBus(g.EventLogger)
	return eng, c
}

// CommitTx is mock
func (g *Gateway) CommitTx(c context.Context) error {
	select {
	case <-c.Done():
		return errors.New("Tx done")
	default:
	}
	return nil
}

type txKey struct{}
type tx struct {
	sync.Mutex
	done bool
}

func (t *tx) Commit() error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	if t.done {
		return errors.New("DONE")
	}
	t.done = true
	fmt.Println("COMMIT!!")
	return nil
}
func (t *tx) Abort() error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	if t.done {
		return errors.New("DONE")
	}
	t.done = true
	fmt.Println("ABORT")
	return nil
}
