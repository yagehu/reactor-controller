package consul

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"

	consulcontroller "github.com/yagehu/reactor-controller/internal/controller/consul"
)

var Module = fx.Invoke(Register)

type Handler interface {
	WatchEvent()
}

type Params struct {
	fx.In

	Lifecycle        fx.Lifecycle
	Router           *mux.Router
	ConsulController consulcontroller.Controller
}

func Register(p Params) {
	h := handler{
		consulController: p.ConsulController,
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			p.Router.HandleFunc("/", h.WatchEvent).Methods(http.MethodPost)

			return nil
		},
		OnStop: nil,
	})
}

type handler struct {
	consulController consulcontroller.Controller
}
