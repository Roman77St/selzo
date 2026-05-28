package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Roman77St/selzo/internal/config"
)

// DB wraps pgxpool.Pool and provides transaction helpers
// and future database-level abstractions.
type DB struct {
	pool *pgxpool.Pool
}

// NewPostgresDB creates a new PostgreSQL connection pool,
// verifies connectivity using Ping, and returns a DB wrapper.
func NewPostgresDB(
	ctx context.Context,
	cfg *config.Config,
) (*DB, error) {

	dsn := fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	cfg.DBHost,
	cfg.DBPort,
	cfg.DBUser,
	cfg.DBPassword,
	cfg.DBName,
	cfg.DBSSLMode,
)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		return nil, fmt.Errorf("ping postgres pool: %w", err)
	}

	return &DB{pool: pool}, nil
}


type TransactionFunc func(tx pgx.Tx) error

// WithTransaction executes fn inside a database transaction.
//
// If fn returns an error, the transaction is rolled back.
// If fn panics, the transaction is rolled back and the panic is rethrown.
// If fn succeeds, the transaction is committed.
func (db *DB) WithTransaction(
	ctx context.Context,
	fn TransactionFunc,
) error {

	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

// Close gracefully closes the PostgreSQL connection pool.
func (db *DB) Close() {
	db.pool.Close()
}