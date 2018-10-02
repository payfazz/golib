package sql

import (
	"context"
	"database/sql"
)

type key int

const queryableKey key = 0

// Queryable ...
type Queryable interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// NewContext ...
func NewContext(ctx context.Context, q Queryable) context.Context {
	return context.WithValue(ctx, queryableKey, q)
}

// FromContext ...
func FromContext(ctx context.Context) (Queryable, bool) {
	q, ok := ctx.Value(queryableKey).(Queryable)
	return q, ok
}
