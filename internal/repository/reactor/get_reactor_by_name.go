package reactor

import (
	"context"

	"github.com/yagehu/reactor-controller/internal/model"
)

type GetReactorByNameParams struct {
	Name string
}

type GetReactorByNameResult struct {
	Record model.Reactor
}

func (r *repository) GetReactorByName(
	ctx context.Context, p *GetReactorByNameParams,
) (*GetReactorByNameResult, error) {
	var record model.Reactor

	err := r.db.QueryRowContext(ctx, `
SELECT
    reactor.id,
    reactor.name,
    reagent.name,
    reactor.created_at,
    reactor.deleted_at
FROM reactor
LEFT JOIN reagent ON reactor.reagent_id = reagent.id
WHERE
	deleted_at IS NOT NULL
LIMIT 1
	`).
		Scan(
			&record.ID,
			&record.Name,
			&record.Reagent.Name,
			&record.CreatedAt,
			&record.DeletedAt,
		)
	if err != nil {
		return nil, err
	}

	return &GetReactorByNameResult{Record: record}, nil
}
