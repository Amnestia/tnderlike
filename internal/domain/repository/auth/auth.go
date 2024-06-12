package authrepo

import (
	"context"
	"database/sql"

	authmodel "github.com/amnestia/tnderlike/internal/domain/model/auth"
	"github.com/amnestia/tnderlike/internal/domain/repository"
	"github.com/amnestia/tnderlike/pkg/logger"
	"github.com/jmoiron/sqlx"
)

// Repository struct
type Repository struct {
	repository.Repository
}

// Auth get credentials from database
func (repo *Repository) Auth(ctx context.Context, email string) (*authmodel.Account, error) {
	acc := &authmodel.Account{}
	err := repo.DB.Slave.QueryRowxContext(ctx, auth, email).Scan(&acc.Email, &acc.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return acc, err
		}
		return acc, logger.ErrorWrap(err, "repo", "Failed on getting account")
	}
	return acc, nil
}

// RegisterNewAccount register new account to database
func (repo *Repository) RegisterNewAccount(ctx context.Context, tx *sqlx.Tx, acc *authmodel.Account) (id int64, err error) {
	q, args, err := sqlx.Named(insertNewAccount, &acc)
	if err != nil {
		return -1, logger.ErrorWrap(err, "repo", "Failed on parsing named parameters")
	}
	q = sqlx.Rebind(sqlx.DOLLAR, q)
	err = tx.GetContext(ctx, &id, q, args...)
	if err != nil {
		return -1, logger.ErrorWrap(err, "repo", "Failed on creating new account")
	}
	return
}
