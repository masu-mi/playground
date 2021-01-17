package dummy

import (
	"testing"

	"github.com/masu-mi/playground/training-code-design/cheap-monster-hunter/game/domain/domaintest"
)

func TestDummyHunterRepository(t *testing.T) {
	domaintest.DoHunterRepositoriesSemanticCheck(t, NewHunterRepo())
}
