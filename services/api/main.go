package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/inn4sc/go-skeleton/config"
	"github.com/inn4sc/go-skeleton/services/api/handler"
	"github.com/inn4sc/vcg-go-common/log"
	"github.com/inn4sc/vcg-go-common/routines"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ctx    context.Context
	logger *logrus.Entry
}

func (s *Server) Init(parentCtx context.Context) routines.Worker {
	return &Server{
		ctx:    parentCtx,
		logger: log.Default.WithField("service", "api-server"),
	}
}

func (s *Server) Run() {
	cfg := config.Cfg{}
	router := GetRouter()
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	server := &http.Server{Addr: addr, Handler: router}

	go func() {
		s.logger.Info("Starting API Server at: ", addr)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.WithError(err).Error("server failed")
		}
	}()

	<-s.ctx.Done()
	s.logger.Info("Shutting down the API Server...")
	serverCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(serverCtx)
	s.logger.Info("Api Server gracefully stopped")
}

func GetRouter() chi.Router {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/dev", func(r chi.Router) {

		r.Route("/ping", func(r chi.Router) {
			r.Get("/", handler.Post)
			r.Get("/buzz/{id}", handler.Post)
			r.Post("/", handler.Post)
		})
	})

	return r
}
