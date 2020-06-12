package handler

import (
	"go.uber.org/fx"

	"github.com/yagehu/reactor-controller/internal/handler/consul"
)

var Module = fx.Options(
	consul.Module,
)
