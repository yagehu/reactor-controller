package mapper

import (
	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/internal/model"
)

func ToReactorModel(x entity.Reactor) model.Reactor {
	return model.Reactor{
		ID:      x.ID.String(),
		Reagent: ToReagentModel(x.Reagent),
	}
}
