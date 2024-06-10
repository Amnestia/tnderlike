package controller

import (
	"github.com/amnestia/tnderlike/internal/domain/controller/auth"
	"github.com/amnestia/tnderlike/internal/domain/controller/ping"
)

// Controller controller containing handler for services
type Controller struct {
	PingHandler ping.Controller
	AuthHandler auth.Controller
}
