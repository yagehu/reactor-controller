package mapper

import (
	"github.com/gofrs/uuid"

	"github.com/yagehu/reactor-controller/internal/entity"
	"github.com/yagehu/reactor-controller/internal/model"
)

func FromReactorModel(x model.Reactor) (entity.Reactor, error) {
	id, err := uuid.FromString(x.ID)
	if err != nil {
		return entity.Reactor{}, err
	}

	reagent, err := FromReagentModule(x.Reagent)
	if err != nil {
		return entity.Reactor{}, err
	}

	return entity.Reactor{
		ID:        id,
		Name:      x.Name,
		Reagent:   reagent,
		CreatedAt: x.CreatedAt,
	}, nil
}

func ToReactorModel(x entity.Reactor) model.Reactor {
	return model.Reactor{
		ID:        x.ID.String(),
		Name:      x.Name,
		Reagent:   ToReagentModel(x.Reagent),
		CreatedAt: x.CreatedAt,
		DeletedAt: nil,
	}
}
