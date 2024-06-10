package authmodel

import "github.com/amnestia/tnderlike/internal/domain/model/common"

// LoginResponse response returned on login
type LoginResponse struct {
	TokenData
	common.DefaultResponse
}
