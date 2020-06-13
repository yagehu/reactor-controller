package controller

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor-controller/internal/controller/reactor"
	"github.com/yagehu/reactor-controller/internal/controller/reactorcontroller"
)

var Module = fx.Provide(
	reactor.New,
	reactorcontroller.New,
)
