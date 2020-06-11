package entity

import (
	"github.com/yagehu/reactor-controller/pkg/apis/reactor/v1alpha1"
)

type EventType string

const (
	EventTypeCreate EventType = "CREATE"
	EventTypeUpdate           = "UPDATE"
	EventTypeDelete           = "DELETE"
)

type Event struct {
	Key         string
	Type        EventType
	ReactorSpec v1alpha1.ReactorSpec
}
