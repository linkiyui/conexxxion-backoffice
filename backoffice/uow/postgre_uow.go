package uow

import (
	"context"
	"errors"
	"time"

	clog "gitlab.com/conexxxion/conexxxion-backoffice/logger"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUnitOfWork struct {
	pgPool *pgxpool.Pool
	pgConn *pgxpool.Conn
	tx     pgx.Tx
}

func NewPostgresUOW(pgPool *pgxpool.Pool) *PostgresUnitOfWork {
	return &PostgresUnitOfWork{
		pgPool: pgPool,
		tx:     nil,
	}
}

func (pguow *PostgresUnitOfWork) BeginTransaction(ctx context.Context) error {
	var err error
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer ctxCancel()
	if pguow.pgConn, err = pguow.pgPool.Acquire(ctx); err != nil {
		clog.ErrorCtx(ctx, "acquiring connection: "+err.Error(), nil)
		return err
	}
	pguow.tx, err = pguow.pgConn.Begin(context.Background())
	if err != nil {
		clog.ErrorCtx(ctx, "starting transaction: "+err.Error(), nil)
	}
	return err
}

func (pguow *PostgresUnitOfWork) Commit(ctx context.Context) error {
	var err error
	defer pguow.pgConn.Release()
	if pguow.tx != nil {
		err = pguow.tx.Commit(context.Background())
		pguow.tx = nil
		if err == nil || errors.Is(err, pgx.ErrTxClosed) {
			return nil
		} else {
			clog.ErrorCtx(ctx, err.Error(), nil)
			return err
		}
	}
	return nil
}

func (pguow *PostgresUnitOfWork) Rollback(ctx context.Context) error {
	var err error
	defer pguow.pgConn.Release()
	if pguow.tx != nil {
		err = pguow.tx.Rollback(context.Background())
		if err == nil || errors.Is(err, pgx.ErrTxClosed) {
			return nil
		} else {
			clog.ErrorCtx(ctx, err.Error(), nil)
			return err
		}
	}
	return nil
}

func (pguow *PostgresUnitOfWork) GetConnection() interface{} {
	return pguow.pgConn
}
