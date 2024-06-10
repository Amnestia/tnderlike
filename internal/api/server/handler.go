package server

import (
	"github.com/amnestia/tnderlike/internal/config"
	"github.com/amnestia/tnderlike/internal/database"
	"github.com/amnestia/tnderlike/internal/domain/controller"
	"github.com/amnestia/tnderlike/internal/domain/controller/auth"
	"github.com/amnestia/tnderlike/internal/domain/controller/ping"
	authrepo "github.com/amnestia/tnderlike/internal/domain/repository/auth"
	authsvc "github.com/amnestia/tnderlike/internal/domain/service/auth"
	pingsvc "github.com/amnestia/tnderlike/internal/domain/service/ping"
	"github.com/amnestia/tnderlike/internal/lib/paseto"
)

type dependency struct {
	db *database.Base

	past *paseto.PASTHandle
}

func getController(cfg config.Config, dep dependency) *controller.Controller {
	pingSvc := &pingsvc.Service{}
	pingCtrl := ping.Controller{
		PingSvc: pingSvc,
	}

	// initialize repo
	authRepo := &authrepo.Repository{}
	authRepo.DB = dep.db

	// initialize service
	authSvc := &authsvc.Service{
		Config: cfg,
		Repo:   authRepo,
		Paseto: dep.past,
	}

	// initialize controller
	authCtrl := auth.Controller{
		AuthSvc: authSvc,
	}

	ctrl := controller.Controller{
		PingHandler: pingCtrl,
		AuthHandler: authCtrl,
	}

	return &ctrl
}
