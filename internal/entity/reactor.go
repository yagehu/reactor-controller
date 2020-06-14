package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Reactor struct {
	ID        uuid.UUID
	Name      string
	IDPrefix  string
	Reagent   Reagent
	CreatedAt time.Time
}
