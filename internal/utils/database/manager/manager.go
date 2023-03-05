package manager

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Manager struct {
	db *gorm.DB
}

func NewManager(db *gorm.DB) *Manager {
	return &Manager{
		db: db,
	}
}

func (m *Manager) WithTransaction(ctx context.Context, f func(ctx context.Context) error) error {
	tx := m.db.Begin()

	ctx = NewTransactionContext(ctx, tx)
	err := f(ctx)
	if err != nil {
		if tErr := tx.Rollback(); tErr != nil {
			return fmt.Errorf("rollback error: %v, original error: %v", tErr, err)
		}
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return fmt.Errorf("error when committing transaction: %v", err)
	}

	return nil
}
