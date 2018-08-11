package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"gitlab.inn4science.com/vcg/go-common/api/render"
	"gitlab.inn4science.com/vcg/go-common/log"
	"gitlab.inn4science.com/vcg/go-common/routines"

	"gitlab.inn4science.com/internal/service-scaffold/config"
	"gitlab.inn4science.com/internal/service-scaffold/workers/api/handler"
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
	router := GetRouter(s.logger, cfg.EnableCORS, cfg.DevMode, cfg.ApiRequestTimeout)
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

func GetRouter(logger *logrus.Entry, enableCORS, enablePPROF bool, requestTimeout int) chi.Router {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	if enableCORS {
		corsHandler := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "jwt", "X-UID"},
			ExposedHeaders:   []string{"Link", "Content-Length"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		})
		r.Use(corsHandler.Handler)
	}

	if enablePPROF {
		logger.Info("API profiler unavailable. Sorry.")
		//r.Mount("/debug", middleware.Profiler())
	}

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	if requestTimeout > 0 {
		t := time.Duration(requestTimeout)
		r.Use(middleware.Timeout(t * time.Second))
	}

	r.Route("/dev", func(r chi.Router) {

		r.Route("/ping", func(r chi.Router) {
			r.Get("/", handler.Post)
			r.Get("/buzz/{id}", handler.Post)
			r.Post("/", handler.Post)
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.ResultNotFound.Render(w)
	})

	return r
}
