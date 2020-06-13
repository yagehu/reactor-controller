package reactor

import (
	"context"

	"github.com/yagehu/reactor-controller/internal/model"
)

type CreateReactorParams struct {
	Reactor model.Reactor
}

type CreateReactorResult struct {
}

func (r *repository) CreateReactor(
	ctx context.Context, p *CreateReactorParams,
) (*CreateReactorResult, error) {
	stmt, err := r.db.Prepare(`
WITH reagent AS (
    INSERT INTO reagent (id, name)
    VALUES ($1, $2)
    RETURNING id
)
INSERT INTO reactor (id, name, reagent_id)
VALUES (
    $3,
    $4,
    (SELECT id FROM reagent)
)
	`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.ExecContext(
		ctx,
		p.Reactor.Reagent.ID,
		p.Reactor.Reagent.Name,
		p.Reactor.ID,
		p.Reactor.Name,
	)
	if err != nil {
		return nil, err
	}

	return &CreateReactorResult{}, nil
}
