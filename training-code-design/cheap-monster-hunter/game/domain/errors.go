package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type ErrNotFound struct{ ID uuid.UUID }

func (e *ErrNotFound) Error() string { return fmt.Sprintf("not found %s", e.ID.String()) }
