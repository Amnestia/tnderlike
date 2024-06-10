package authsvc

import (
	"github.com/amnestia/tnderlike/internal/config"
	"github.com/amnestia/tnderlike/internal/domain/repository"
	"github.com/amnestia/tnderlike/internal/lib/paseto"
)

// Service service functionality of the domain
type Service struct {
	Config config.Config
	Repo   repository.AuthRepository
	Paseto paseto.Handler
}
