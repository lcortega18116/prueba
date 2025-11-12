package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
)

func OpenCockroach(dsn string) (*sqlx.DB, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.MaxConnLifetime = 60 * time.Minute
	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}
	return sqlx.NewDb(pool.StdConn(), "pgx"), nil
}

// Reintentos para SERIALIZATION_FAILURE (SQLSTATE 40001)
func WithTxRetry(ctx context.Context, db *sqlx.DB, fn func(*sqlx.Tx) error) error {
	for i := 0; i < 5; i++ {
		tx, err := db.BeginTxx(ctx, nil)
		if err != nil {
			return err
		}
		if err = fn(tx); err != nil {
			_ = tx.Rollback()
			var pqe interface{ SQLState() string }
			if errors.As(err, &pqe) && pqe.SQLState() == "40001" {
				time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
				continue
			}
			return err
		}
		if err = tx.Commit(); err != nil {
			var pqe interface{ SQLState() string }
			if errors.As(err, &pqe) && pqe.SQLState() == "40001" {
				time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
				continue
			}
			return err
		}
		return nil
	}
	return errors.New("transaction failed after retries")
}
