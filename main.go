package main

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor-controller/config"
	"github.com/yagehu/reactor-controller/internal/controller"
	"github.com/yagehu/reactor-controller/internal/fx/httpfx"
	"github.com/yagehu/reactor-controller/internal/fx/kubernetesfx"
	"github.com/yagehu/reactor-controller/internal/fx/middlewarefx"
	"github.com/yagehu/reactor-controller/internal/fx/zapfx"
	"github.com/yagehu/reactor-controller/internal/handler"
	"github.com/yagehu/reactor-controller/internal/reactor"
)

func main() {
	fx.New(opts()).Run()
}

func opts() fx.Option {
	return fx.Options(
		config.Module,
		controller.Module,
		handler.Module,
		reactor.Module,

		httpfx.Module,
		middlewarefx.Module,
		kubernetesfx.Module,
		zapfx.Module,
	)
}
