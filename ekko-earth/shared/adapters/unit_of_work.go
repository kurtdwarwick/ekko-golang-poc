package adapters

import "context"

type UnitOfWork interface {
	Execute(
		execution func(transaction Transaction, context context.Context) (any, error),
		context context.Context,
	) (any, error)
}
