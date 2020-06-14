package mapper

import (
	"github.com/gofrs/uuid"

	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/internal/model"
)

func FromReagentModel(x model.Reagent) (entity.Reagent, error) {
	id, err := uuid.FromString(x.ID)
	if err != nil {
		return entity.Reagent{}, err
	}

	return entity.Reagent{
		ID:       id,
		Name:     x.Name,
		IDPrefix: x.IDPrefix,
	}, nil
}

func ToReagentModel(x entity.Reagent) model.Reagent {
	return model.Reagent{
		ID:       x.ID.String(),
		Name:     x.Name,
		IDPrefix: x.IDPrefix,
	}
}
