package zapfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/yagehu/reactor-controller/config"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Config config.Config
}

type Result struct {
	fx.Out

	Logger *zap.Logger
}

func New(p Params) (Result, error) {
	var (
		logger *zap.Logger
		err    error
	)

	switch p.Config.RuntimeEnvironment {
	case config.RuntimeEnvironmentProduction:
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		return Result{}, err
	}

	return Result{
		Logger: logger,
	}, nil
}
