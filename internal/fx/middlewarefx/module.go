package middlewarefx

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor-controller/internal/fx/middlewarefx/logging"
)

var Module = fx.Provide(
	logging.New,
)
