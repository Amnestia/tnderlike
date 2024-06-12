package authsvc

import (
	"github.com/amnestia/tnderlike/internal/lib/paseto"
	mocks "github.com/amnestia/tnderlike/mocks/internal_/domain/repository"
)

var (
	mockRepo        *mocks.AuthRepository
	mockPASTHandler *paseto.MockHandlerPaseto
)

func initMock() {
	mockRepo = new(mocks.AuthRepository)
	mockPASTHandler = new(paseto.MockHandlerPaseto)
}

func getMock() *Service {
	initMock()
	return &Service{
		Paseto: mockPASTHandler,
		Repo:   mockRepo,
	}
}
