package adapters

type Transaction interface {
	Begin() any
	Commit() error
	Rollback() error
}
