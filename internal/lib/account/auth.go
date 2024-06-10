package account

import (
	"context"
	"errors"

	"github.com/amnestia/tnderlike/internal/lib/paseto"
	"github.com/amnestia/tnderlike/pkg/logger"
)

// Data struct containing account data from token
type Data struct {
	ID       int64
	Username string
	Email    string
}

// GetData get account data from context inserted from middleware
func GetData(ctx context.Context) (*Data, error) {
	payload := ctx.Value(paseto.AuthData)
	if payload == nil {
		return nil, logger.ErrorWrap(errors.New("nil on context"), "auth", "invalid data")
	}
	p := payload.(paseto.Payload)
	return &Data{
		ID:       p.ID,
		Username: p.Username,
		Email:    p.Email,
	}, nil
}
