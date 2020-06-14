package reactor

import (
	"context"
	"database/sql"

	"github.com/yagehu/reactor-controller/internal/errs"
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
	var (
		op     errs.Op = "repository/reactor.GetReactorByName"
		record model.Reactor
	)

	err := r.db.QueryRowContext(
		ctx,
		`
SELECT
    reactor.id,
    reactor.name,
    reagent.id,
    reagent.name,
    reagent.id_prefix,
    reactor.created_at,
    reactor.deleted_at
FROM reactor
LEFT JOIN reagent ON reactor.reagent_id = reagent.id
WHERE
    reactor.name = $1
    AND deleted_at IS NULL
LIMIT 1
		`,
		p.Name,
	).
		Scan(
			&record.ID,
			&record.Name,
			&record.Reagent.ID,
			&record.Reagent.Name,
			&record.Reagent.IDPrefix,
			&record.CreatedAt,
			&record.DeletedAt,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.E(op, errs.KindReactorNotFound, err)
		}

		return nil, errs.E(op, err)
	}

	return &GetReactorByNameResult{Record: record}, nil
}
