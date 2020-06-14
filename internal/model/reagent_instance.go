package model

import "time"

type ReagentInstance struct {
	ID        string
	Namespace string
	Name      string
	CreatedAt time.Time
}
