package repository

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor-controller/internal/repository/reactor"
)

var Module = fx.Provide(reactor.New)
