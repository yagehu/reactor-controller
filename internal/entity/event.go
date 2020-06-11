package entity

import (
	"github.com/yagehu/reactor-controller/pkg/apis/reactor/v1alpha1"
)

type Event struct {
	Key         string
	ReactorSpec v1alpha1.ReactorSpec
}
