package repository

import (
	"context"

	"github.com/amnestia/tnderlike/internal/database"
	"github.com/amnestia/tnderlike/pkg/logger"
	"github.com/jmoiron/sqlx"
)

// Repository struct containing repository databases
type Repository struct {
	DB *database.Base
}

// NewTransaction new transaction
func (repo *Repository) NewTransaction(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := repo.DB.Master.BeginTxx(ctx, nil)
	if err != nil {
		return nil, logger.ErrorWrap(err, "database", "Error on begin transaction")
	}
	return tx, nil
}

// Commit commit transaction
func (repo *Repository) Commit(tx *sqlx.Tx) (err error) {
	if err = tx.Commit(); err != nil {
		return logger.ErrorWrap(err, "database", "Error on commit")
	}
	return nil
}

// RollbackOnError rollback transaction if error occurs
func (repo *Repository) RollbackOnError(tx *sqlx.Tx, err error) error {
	tmpErr := err
	if err == nil {
		return nil
	}
	if err = tx.Rollback(); err != nil {
		return logger.ErrorWrap(err, "database", "Error on rollback")
	}

	return tmpErr
}
