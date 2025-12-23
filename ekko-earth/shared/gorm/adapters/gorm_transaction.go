package adapters

import "gorm.io/gorm"

type GormTransaction struct {
	Transaction *gorm.DB
}

func NewGormTransaction(transaction *gorm.DB) *GormTransaction {
	return &GormTransaction{Transaction: transaction}
}

func (transaction *GormTransaction) Begin() any {
	return transaction.Transaction.Begin()
}

func (transaction *GormTransaction) Commit() error {
	transaction.Transaction.Commit()

	return nil
}

func (transaction *GormTransaction) Rollback() error {
	transaction.Transaction.Rollback()

	return nil
}
