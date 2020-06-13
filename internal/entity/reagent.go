package entity

import (
	"github.com/gofrs/uuid"
)

type Reagent struct {
	ID   uuid.UUID
	Name string
}
