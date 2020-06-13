package reactorcontroller

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"

	reactorcontroller "github.com/yagehu/reactor-controller/internal/controller/reactor"
	"github.com/yagehu/reactor-controller/internal/errs"
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
	reactor, err := c.reactorLister.Reactors(p.Namespace).Get(p.Name)
	if err != nil {
		// The Foo resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			_, err := c.reactorController.DeleteReactor(
				ctx, &reactorcontroller.DeleteReactorParams{Name: p.Name},
			)
			if err != nil {
				if e, ok := err.(*errs.Error); ok {
					if e.Kind == errs.KindReactorNotFound {
						return &ProcessEventResult{}, nil
					}
				}

				return nil, err
			}

			return &ProcessEventResult{}, nil
		}

		return nil, err
	}

	_, err = c.reactorController.GetReactorByName(
		ctx, &reactorcontroller.GetReactorByNameParams{Name: reactor.Name},
	)
	if err != nil {
		e, ok := err.(*errs.Error)
		if !ok {
			return nil, err
		}

		if e.Kind != errs.KindReactorNotFound {
			return nil, err
		}

		_, err := c.reactorController.CreateReactor(
			ctx, &reactorcontroller.CreateReactorParams{
				Name:        reactor.Name,
				ReagentName: reactor.Spec.Reagent.Name,
			},
		)
		if err != nil {
			return nil, err
		}
	}

	return &ProcessEventResult{}, nil
}
