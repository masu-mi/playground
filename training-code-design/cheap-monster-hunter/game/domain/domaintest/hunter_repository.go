package domaintest

import (
	"context"
	"errors"
	"testing"

	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain"
)

func DoHunterRepositoriesSemanticCheck(t *testing.T, repo domain.HunterRepository) {
	var inputs []*domain.Hunter
	for _, name := range []string{"apple", "orange"} {
		h, _ := domain.NewStandardHunter(name)
		inputs = append(inputs, h)
	}
	t.Run("Save() and FindByID() succeed", func(t *testing.T) {
		for _, i := range inputs {
			e := repo.Save(context.Background(), i)
			if e != nil {
				t.Errorf("Save failed with %v", i)
			}

			got, e := repo.FindByID(context.Background(), i.ID)
			if e != nil {
				t.Errorf("FindByID failed with %v", i)
			}
			if got.ID != i.ID {
				t.Errorf("FindByID returns a hunter (%v) don't match with passed ID (%s)", got, i.ID)
			}
		}
	})

	t.Run("Remove() work well", func(t *testing.T) {
		target, other := inputs[0], inputs[1]
		e := repo.Remove(context.Background(), target)
		if e != nil {
			t.Errorf("Remove(%v) failed", target)
		}

		var errNotFound *domain.ErrNotFound
		_, e = repo.FindByID(context.Background(), target.ID)
		if !errors.As(e, &errNotFound) {
			t.Errorf("Found or not expected error: %v", e)
		}
		_, e = repo.FindByID(context.Background(), other.ID)
		if e != nil {
			t.Errorf("Fail to FindByID() on un removed target: err = %v", e)
		}
	})
}
