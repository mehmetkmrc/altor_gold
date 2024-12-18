package psql

import (
	"context"
	
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/config"
	"github.com/mehmetkmrc/ator_gold/internal/core/port/db"
	"go.uber.org/zap"
)


var (
	_ db.EngineMaker = (*pdb)(nil)
)

type (
	pdb struct{
		cfg		*config.Container
		queryBuilder *squirrel.StatementBuilderType
		pool 		*pgxpool.Pool
	}

)

func NewDB(cfg *config.Container) db.EngineMaker {
	psqlDB := &pdb{
		cfg: cfg,
	}
	queryBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	psqlDB.queryBuilder = &queryBuilder

	return psqlDB
}

func (ps *pdb) Start(ctx context.Context) error {
	url := ps.getURL()
	err := ps.connect(ctx, url)
	if err != nil {
		zap.S().Fatal("Postgres connection failed", err)
	}

	zap.S().Info("Connected to Postgres 🎉")
	return nil
}

func (ps *pdb) getURL() string {
	url := ps.cfg.PSQL.URL
	return url
}

func (ps *pdb) ping(ctx context.Context) error {
	if ps.pool != nil {
		if err := ps.pool.Ping(ctx); err != nil {
			return err
		}
	}
	zap.S().Info("Postgres is ready to serve")
	return nil
}

func (ps *pdb) connect(ctx context.Context, url string) error {
	var lastErr error
	for ps.cfg.Settings.PSQLConnAttempts > 0 {
		zap.S().Info("Connecting to Postgres...")
		ps.pool, lastErr = pgxpool.New(ctx, url)
		if lastErr == nil {
			err := ps.ping(ctx)
			if err == nil {
				zap.S().Info("Postgres Pong! 🐘")
				return nil	
			}
		}

		ps.cfg.Settings.PSQLConnAttempts--
		zap.S().Warnf("Postgres connection failed, attempts left: %d", ps.cfg.Settings.PSQLConnAttempts)
		time.Sleep(time.Duration(ps.cfg.Settings.PSQLConnTimeout) * time.Second)
	}

	panic("Postgres connection failed")
}



func (ps *pdb) Close(ctx context.Context) error {
	zap.S().Info("Postgres Context is done. Shutting down server...")
	ps.pool.Close()
	return nil
}

func (ps *pdb) GetDB() *pgxpool.Pool {
	return ps.pool
}

func (ps *pdb) Execute(ctx context.Context, query string, args ...any) error {
	_, err := ps.pool.Exec(ctx, query, args...)
	return err
}

func (ps *pdb) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return ps.pool.Query(ctx, sql, args...)
}

func (ps *pdb) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return ps.pool.QueryRow(ctx, sql, args...)
}
