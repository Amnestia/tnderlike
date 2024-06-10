package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/amnestia/tnderlike/internal/api/server/router"
	"github.com/amnestia/tnderlike/internal/config"
	"github.com/amnestia/tnderlike/internal/database"
	"github.com/amnestia/tnderlike/internal/lib/paseto"
	"github.com/amnestia/tnderlike/internal/middleware/auth"
	"github.com/amnestia/tnderlike/pkg/logger"
)

// Server struct containing server config and options
type Server struct {
	Cfg    config.Config
	Router *router.Router
}

// New initialize server
func New() *Server {
	var err error

	// get config
	cfg := config.Config{}
	cfg = cfg.ReadJSONConfig("server")
	cfg = cfg.ReadYAMLConfig("server")
	err = logger.InitLogger(cfg.App, cfg.Server.Logs.Info, cfg.Server.Logs.Error)
	if err != nil {
		log.Fatal("Error on creating log files : ", err)
		return nil
	}

	// initialize db, etc
	dep := dependency{}
	dep.db, err = database.New(cfg.Database, "postgres")
	if err != nil {
		log.Fatal("Error on connecting to database : ", err)
		return nil
	}

	controller := getController(cfg, dep)

	// initialize interactor(controller, service)
	dep.past = paseto.New(cfg)
	auth := auth.AuthorizationModule{
		Token: dep.past,
	}
	// initialize router
	r := router.New(router.Options{}, cfg, auth, controller)
	return &Server{
		Cfg:    cfg,
		Router: r,
	}
}

// Run run server
func (s *Server) Run() int {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", s.Cfg.Server.Port),
		Handler:      s.Router.Handler,
		ReadTimeout:  time.Duration(s.Cfg.Server.Timeout) * time.Second,
		WriteTimeout: time.Duration(s.Cfg.Server.Timeout) * time.Second,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		go func() {
			<-ctx.Done()
			if ctx.Err() == context.DeadlineExceeded {
				log.Fatal("Graceful Shutdown timed out")
			}
		}()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal(logger.ErrorWrap(err, "server.Shutdown"))
		}
	}()
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(logger.ErrorWrap(err, "server.ListenAndServe"))
		return 1
	}
	return 0
}
