package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EngineMaker interface{
	Start(ctx context.Context) error
	Close(ctx context.Context) error
	GetDB() *pgxpool.Pool
	Execute(ctx context.Context, query string, args ...any) error
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}