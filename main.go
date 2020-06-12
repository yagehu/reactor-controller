package main

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor-controller/config"
	"github.com/yagehu/reactor-controller/internal/controller"
	"github.com/yagehu/reactor-controller/internal/fx/kubernetesfx"
	"github.com/yagehu/reactor-controller/internal/fx/zapfx"
	"github.com/yagehu/reactor-controller/internal/reactorcontroller"
)

func main() {
	fx.New(opts()).Run()
}

func opts() fx.Option {
	return fx.Options(
		config.Module,
		controller.Module,
		reactorcontroller.Module,

		kubernetesfx.Module,
		zapfx.Module,
	)
}
