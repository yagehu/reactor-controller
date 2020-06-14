package consul

import (
	"encoding/json"
	"net/http"

	consulcontroller "github.com/yagehu/reactor-controller/internal/controller/consul"
	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/internal/httprespond"
)

func (h *handler) WatchEvent(w http.ResponseWriter, r *http.Request) {
	var request WatchEventRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		httprespond.WithError(
			w, http.StatusBadRequest, httprespond.MsgCouldNotDecodeRequest,
		)

		return
	}

	reagentInstances := make([]entity.Source, 0, len(request))

	for name, tags := range request {
		tagsMap := make(map[string]struct{})

		for _, tag := range tags {
			tagsMap[tag] = struct{}{}
		}

		reagentInstances = append(reagentInstances, entity.Source{
			Name: name,
			Tags: tagsMap,
		})
	}

	_, err := h.consulController.HandleWatchEvent(
		r.Context(),
		&consulcontroller.HandleWatchEventParams{
			Sources: reagentInstances,
		},
	)
	if err != nil {
		httprespond.WithError(
			w,
			http.StatusInternalServerError,
			httprespond.MsgCouldNotHandleWatchEvent,
		)

		return
	}

	httprespond.With(w, WatchEventResponse{})
}
