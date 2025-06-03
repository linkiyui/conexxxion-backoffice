package infra

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/uow"
)

type PostgreRepository struct {
	pgPool     *pgxpool.Pool
	pgPoolConn *pgxpool.Conn
	uow        uow.UnitOfWork
}

func NewPostgreRepository(pgPool *pgxpool.Pool) *PostgreRepository {
	return &PostgreRepository{
		pgPool:     pgPool,
		pgPoolConn: nil,
		uow:        nil,
	}
}

func (r *PostgreRepository) WithUnitOfWork(uw uow.UnitOfWork) *PostgreRepository {
	r.uow = uw
	return r
}

func (r *PostgreRepository) GetConn() (*pgxpool.Conn, error) {
	var err error
	var ok bool
	if r.uow != nil {
		if r.pgPoolConn, ok = r.uow.GetConnection().(*pgxpool.Conn); ok {
			return r.pgPoolConn, nil
		}
		return nil, errors.New("cannot parse transaction from unit of work")
	}
	var ctxCancel context.CancelFunc

	ctx := context.Background()
	ctx, ctxCancel = context.WithTimeout(ctx, time.Second*5)
	defer ctxCancel()
	r.pgPoolConn, err = r.pgPool.Acquire(ctx)
	if err != nil {
		return nil, errors.Join(errors.New("cannot acquire connection from pool"), err)
	}
	return r.pgPoolConn, nil
}

func (r *PostgreRepository) ReleaseConn() {
	if r.uow == nil {
		r.pgPoolConn.Release()
	}
}
