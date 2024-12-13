package db

import (
	"context"
	"database/sql"
)

type EngineMaker interface{
	Start(ctx context.Context) error
	Close(ctx context.Context) error
	GetDB() *sql.DB
	Execute(query string, args ...any) error
	Query(sql string, args ...any) (*sql.Rows, error)
	QueryRow(sql string, args ...any) *sql.Row
}