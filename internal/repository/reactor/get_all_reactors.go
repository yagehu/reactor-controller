package reactor

import (
	"context"

	"go.uber.org/zap"

	"github.com/yagehu/reactor-controller/internal/errs"
	"github.com/yagehu/reactor-controller/internal/model"
)

type GetAllReactorsParams struct {
}

type GetAllReactorsResult struct {
	Records []model.Reactor
}

func (r *repository) GetAllReactors(
	ctx context.Context, p *GetAllReactorsParams,
) (*GetAllReactorsResult, error) {
	var (
		op      errs.Op = "repository/reactor.GetAllReactors"
		records []model.Reactor
	)

	rows, err := r.db.QueryContext(ctx, `
SELECT
    reactor.id,
    reactor.name,
    reagent.id,
    reagent.name,
    reactor.created_at
FROM
    reactor
LEFT JOIN
    reagent ON reactor.reagent_id = reagent.id
WHERE
    deleted_at IS NULL
	`)
	if err != nil {
		return nil, errs.E(op, err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			r.logger.Error("Could not close rows.", zap.Error(err))
		}
	}()

	for rows.Next() {
		var reactor model.Reactor

		err := rows.Scan(
			&reactor.ID,
			&reactor.Name,
			&reactor.Reagent.ID,
			&reactor.Reagent.Name,
			&reactor.CreatedAt,
		)
		if err != nil {
			return nil, errs.E(op, err)
		}

		records = append(records, reactor)
	}

	return &GetAllReactorsResult{
		Records: records,
	}, nil
}
