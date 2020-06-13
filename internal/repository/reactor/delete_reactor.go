package reactor

import (
	"context"
	"time"
)

type DeleteReactorParams struct {
	ID string
}

type DeleteReactorResult struct {
}

func (r *repository) DeleteReactor(
	ctx context.Context, p *DeleteReactorParams,
) (*DeleteReactorResult, error) {
	stmt, err := r.db.Prepare(`
UPDATE reactor
SET
	deleted_at = $1
WHERE
	id = $2
	AND deleted_at IS NULL
	`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.ExecContext(ctx, time.Now(), p.ID)
	if err != nil {
		return nil, err
	}

	return &DeleteReactorResult{}, nil
}
