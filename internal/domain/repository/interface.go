package repository

import (
	"context"

	authmodel "github.com/amnestia/tnderlike/internal/domain/model/auth"
	"github.com/jmoiron/sqlx"
)

// Repositorier interface
type Repositorier interface {
	NewTransaction(ctx context.Context) (*sqlx.Tx, error)
	Commit(tx *sqlx.Tx) (err error)
	RollbackOnError(tx *sqlx.Tx, err error) error
}

// AuthRepository interface
type AuthRepository interface {
	Repositorier
	Auth(ctx context.Context, email string) (*authmodel.Account, error)
	RegisterNewAccount(ctx context.Context, tx *sqlx.Tx, acc *authmodel.Account) (id int64, err error)
}
