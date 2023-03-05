package manager

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type key string

const dbTransKey key = "db_trans"

type DBTx interface {
	// method from gorm.DB
	WithContext(ctx context.Context) *gorm.DB
	DB() (*sql.DB, error)
}

func NewTransactionContext(ctx context.Context, q DBTx) context.Context {
	return context.WithValue(ctx, dbTransKey, q)
}

func DBTransactionFromContext(ctx context.Context) (DBTx, bool) {
	conn, ok := ctx.Value(dbTransKey).(DBTx)
	return conn, ok
}
