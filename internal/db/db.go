package db

import (
	"context"
	"database/sql"
	"snowApp/internal/model"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
}

type Queries struct {
	db DBTX
}

// CreateUser implements Querier.
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (model.User, error) {
	panic("unimplemented")
}

// GetUserByID implements Querier.
func (q *Queries) GetUserByID(ctx context.Context, id string) (model.User, error) {
	panic("unimplemented")
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}
