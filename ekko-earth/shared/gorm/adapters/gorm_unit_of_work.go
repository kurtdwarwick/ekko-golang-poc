package adapters

import (
	"context"

	"github.com/ekko-earth/shared/adapters"
)

type GormUnitOfWork struct {
	database GormDatabase
}

func NewGormUnitOfWork(database GormDatabase) *GormUnitOfWork {
	return &GormUnitOfWork{database: database}
}

func (unitOfWork *GormUnitOfWork) Execute(
	execution func(transaction adapters.Transaction, ctx context.Context) (any, error),
	ctx context.Context,
) (any, error) {
	transaction := NewGormTransaction(unitOfWork.database.Database.Begin())

	result, err := execution(transaction, ctx)

	if err != nil {
		transaction.Rollback()
		return nil, err
	}

	transaction.Commit()

	return result, nil
}
