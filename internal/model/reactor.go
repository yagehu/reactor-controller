package model

import (
	"time"
)

type Reactor struct {
	ID        string
	Name      string
	Reagent   Reagent
	CreatedAt time.Time
	DeletedAt *time.Time
}
