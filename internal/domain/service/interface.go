package service

import (
	"context"

	authmodel "github.com/amnestia/tnderlike/internal/domain/model/auth"
	"github.com/amnestia/tnderlike/internal/domain/model/common"
)

// PingServicer interface
type PingServicer interface {
	Ping() string
}

// AuthServicer interface
type AuthServicer interface {
	Register(context.Context, *authmodel.Account) *common.DefaultResponse
	Auth(context.Context, *authmodel.Account) *authmodel.LoginResponse
}
