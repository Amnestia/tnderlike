package router

import (
	"github.com/amnestia/tnderlike/internal/config"
	"github.com/amnestia/tnderlike/internal/domain/controller"
	intmiddleware "github.com/amnestia/tnderlike/internal/middleware"
	"github.com/amnestia/tnderlike/internal/middleware/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Router struct containing router wrapper
type Router struct {
	Handler    chi.Router
	Controller *controller.Controller
	Config     config.Config
	Options    Options
	Auth       auth.AuthorizationModule
}

type Options struct{}

func New(opt Options, cfg config.Config, auth auth.AuthorizationModule, controller *controller.Controller) *Router {
	r := &Router{
		Handler:    chi.NewRouter(),
		Options:    opt,
		Config:     cfg,
		Auth:       auth,
		Controller: controller,
	}
	r.setMiddleware()
	return r
}

// setMiddleware set default middleware
func (r *Router) setMiddleware() {
	r.Handler.Use(cors.Handler(cors.Options{
		AllowedOrigins:   r.Config.Server.Origin,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Handler.Use(intmiddleware.PanicRecovery)
	r.Handler.Use(middleware.CleanPath)
	r.Handler.Use(middleware.RealIP)
	r.Handler.Use(middleware.RequestID)
	r.Handler.Use(intmiddleware.Logger)
}

// registerRoute register API routes
func (router *Router) registerRoute() {
	router.Handler.Get("/ping", router.Controller.PingHandler.Ping)

}
