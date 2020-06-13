package entity

import (
	"github.com/gofrs/uuid"
)

type Reactor struct {
	ID      uuid.UUID
	Reagent Reagent
}
