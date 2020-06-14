package entity

import (
	"github.com/gofrs/uuid"
)

type ReagentInstance struct {
	ID       uuid.UUID
	Reactors []Reactor
}
