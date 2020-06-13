package mapper

import (
	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/internal/model"
)

func ToReagentModel(x entity.Reagent) model.Reagent {
	return model.Reagent{
		ID:   x.ID.String(),
		Name: x.Name,
	}
}
