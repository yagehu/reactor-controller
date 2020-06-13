package reactorcontroller

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"

	reactorcontroller "github.com/yagehu/reactor-controller/internal/controller/reactor"
)

type ProcessEventParams struct {
	Name      string
	Namespace string
}

type ProcessEventResult struct {
}

func (c *controller) ProcessEvent(
	ctx context.Context, p *ProcessEventParams,
) (*ProcessEventResult, error) {
	_, err := c.reactorLister.Reactors(p.Namespace).Get(p.Name)
	if err != nil {
		// The Foo resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			_, err := c.reactorController.DeleteReactor(
				ctx, &reactorcontroller.DeleteReactorParams{Name: p.Name},
			)
			if err != nil {
				return nil, err
			}

			return &ProcessEventResult{}, nil
		}

		return nil, err
	}

	return &ProcessEventResult{}, nil
}
