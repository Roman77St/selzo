package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// txKey is a private context key used to store database transactions.
type txKey struct{}

var transactionContextKey = txKey{}

// WithinTransaction выполняет переданную функцию fn внутри транзакции Postgres
func (db *DB) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {

	if _, ok := ctx.Value(transactionContextKey).(pgx.Tx); ok {
		return fn(ctx)
	}
	// Стартуем транзакцию
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	// Гарантируем откат при панике или ошибке
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	// Кладем транзакцию в контекст и передаем в функцию с бизнес-логикой
	txCtx := context.WithValue(ctx, transactionContextKey, tx)
	if err := fn(txCtx); err != nil {
		return err
	}

	// Если всё успешно — коммитим
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func DBTXFromContext(ctx context.Context, defaultDB DBTX) DBTX {
    if tx, ok := ctx.Value(transactionContextKey).(pgx.Tx); ok {
        return tx
    }
    return defaultDB
}