package psql

import (
	"context"
	"database/sql"
	"github.com/mehmetkmrc/ator_gold/internal/adapter/secondary/config"
	"time"
	_ "github.com/lib/pq"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"fmt"
)

type (
	EngineMaker interface {
		Start(ctx context.Context) error
		Close(ctx context.Context) error
		GetDB() *sql.DB
		Execute(query string, args ...any) error
		Query(sql string, args ...any) (*sql.Rows, error)
		QueryRow(sql string, args ...any) *sql.Row
	}

	MSDB struct {
		cfg          *config.Container
		queryBuilder *squirrel.StatementBuilderType
		db           *sql.DB
	}


)

func NewMSDB(cfg *config.Container) *MSDB {
	db := &MSDB{
		cfg: cfg,
	}
	queryBuilder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	db.queryBuilder = &queryBuilder

	return db
}

func (ms *MSDB) Start(ctx context.Context) error {
	url := ms.getURL()
	err := ms.connect(url)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MSDB) getURL() string {
	url := ms.cfg.PSQL.URL
	return url
}

func (ms *MSDB) connect(url string) error {
	var lastErr error
	for ms.cfg.Settings.PSQLConnAttempts > 0 {
		zap.S().Info("Connecting to PSQL...")
		ms.db, lastErr = sql.Open("postgres", url)

		if lastErr == nil {
			err := ms.ping()
			if err == nil {
				zap.S().Info("PSQL Pong!")
				return nil
			} else {
                zap.S().Warnf("PSQL ping failed: %v", err) 
            }
		} else {
            zap.S().Warnf("PSQL connection failed: %v", lastErr)
        }

		ms.cfg.Settings.PSQLConnAttempts--
		zap.S().Warnf("PSQL connection failed, attempts left: %d", ms.cfg.Settings.PSQLConnAttempts)
		time.Sleep(time.Duration(ms.cfg.Settings.PSQLConnTimeout) * time.Second)
	}

	return fmt.Errorf("PSQL connection failed after %d attempts", ms.cfg.Settings.PSQLConnAttempts)
}

func (ms *MSDB) ping() error {
	if ms.db != nil {
		if err := ms.db.Ping(); err != nil {
			return err
		}
	}
	return nil
}

func (ms *MSDB) Close(ctx context.Context) error {
	if ms.db != nil {
		return ms.db.Close()
	}
	return nil
}

func (ms *MSDB) GetDB() *sql.DB {
	return ms.db
}

func (ms *MSDB) Execute(query string, args ...interface{}) error {
	_, err := ms.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}


func (ms *MSDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return ms.db.Query(query, args...)
}

func (ms *MSDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return ms.db.QueryRow(query, args...)
}