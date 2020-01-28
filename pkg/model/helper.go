package model

import (
	"context"
	"database/sql"
)

type DBConnect interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func GetDB2DBConnect(getter func() *sql.DB) func() DBConnect {
	return func() DBConnect {
		return getter()
	}
}

func Tx2DBConnect(tx *sql.Tx) func() DBConnect {
	return func() DBConnect {
		return tx
	}
}
