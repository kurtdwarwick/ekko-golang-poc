package adapters

import "context"

type UnitOfWork interface {
	Execute(
		execution func(transaction Transaction, ctx context.Context) (any, error),
		ctx context.Context,
	) (any, error)
}
