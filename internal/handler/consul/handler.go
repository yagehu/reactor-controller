package consul

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

var Module = fx.Invoke(Register)

type Handler interface {
	WatchEvent()
}

type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Router    *mux.Router
}

func Register(p Params) {
	h := handler{}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			p.Router.HandleFunc("/", h.WatchEvent).Methods(http.MethodPost)

			return nil
		},
		OnStop: nil,
	})
}

type handler struct {
}

func (h *handler) WatchEvent(w http.ResponseWriter, r *http.Request) {

}
