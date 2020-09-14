package sql_test

import (
	"context"
	"database/sql"
	"testing"

	s "github.com/payfazz/golib/pkg/sql"
)

func TestQueryableContext(t *testing.T) {
	q := &queryable{}
	ctx := s.NewContext(context.TODO(), q)

	que, ok := s.FromContext(ctx)

	if !ok {
		t.Error("failed to get Queryable from context")
	}
	if que != q {
		t.Error("queryable is invalid")
	}
}

type queryable struct {
	s.Queryable
}

func (q *queryable) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}
func (q *queryable) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}
func (q *queryable) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}
func (q *queryable) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return nil
}
func (q *queryable) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}
func (q *queryable) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}
