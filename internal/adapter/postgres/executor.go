package postgres

import (
	"context"
	"database/sql"

	txManager "Olegnemlii/wallet-service/internal/adapter/txmanager"
)

type DB interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Executor struct {
	db *sql.DB
}

func NewExecutor(db *sql.DB) *Executor {
	return &Executor{db: db}
}

func (e *Executor) Get(ctx context.Context) DB {
	if tx, ok := ctx.Value(txManager.TxKeyName).(*sql.Tx); ok {
		return tx
	}
	return e.db
}
