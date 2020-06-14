package controller

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor-controller/internal/controller/consul"
	"github.com/yagehu/reactor-controller/internal/controller/reactor"
	"github.com/yagehu/reactor-controller/internal/controller/reactorcontroller"
	"github.com/yagehu/reactor-controller/internal/controller/reagentinstance"
)

var Module = fx.Provide(
	consul.New,
	reactor.New,
	reactorcontroller.New,
	reagentinstance.New,
)
