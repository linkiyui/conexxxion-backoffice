package uow

import "context"

type UnitOfWork interface {
	BeginTransaction(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	GetConnection() interface{}
}
