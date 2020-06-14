package reagentinstance

import (
	"context"

	"github.com/yagehu/reactor-controller/internal/errs"
	"github.com/yagehu/reactor-controller/internal/model"
)

type GetAllReagentInstancesParams struct {
	Namespace string
}

type GetAllReagentInstancesResult struct {
	Records []model.ReagentInstance
}

func (r *repository) GetAllReagentInstances(
	ctx context.Context, p *GetAllReagentInstancesParams,
) (*GetAllReagentInstancesResult, error) {
	var (
		op      errs.Op = "repository/reagentinstance.GetAllReagentInstances"
		records []model.ReagentInstance
	)

	rows, err := r.db.QueryContext(
		ctx,
		`
SELECT
    id,
    namespace,
    name,
    created_at
FROM
    reagent_instance
WHERE
    deleted_at IS NULL
		`,
		p.Namespace,
	)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for rows.Next() {
		var record model.ReagentInstance

		err := rows.Scan(
			&record.ID,
			&record.Namespace,
			&record.Name,
			&record.CreatedAt,
		)
		if err != nil {
			return nil, errs.E(op, err)
		}

		records = append(records, record)
	}

	return &GetAllReagentInstancesResult{Records: records}, nil
}
