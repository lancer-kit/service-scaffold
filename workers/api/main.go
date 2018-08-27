package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"gitlab.inn4science.com/gophers/service-kit/api"
	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/auth"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
	"gitlab.inn4science.com/gophers/service-scaffold/workers/api/handler"
	"gitlab.inn4science.com/gophers/service-scaffold/workers/api/middlewares"
)

func Server() *api.Server {
	return &api.Server{
		Name:      "api-server",
		GetRouter: GetRouter,
		GetConfig: func() api.Config {
			return config.Config().Api
		},
	}
}

func GetRouter(logger *logrus.Entry, config api.Config) http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(log.NewRequestLogger(logger.Logger))
	r.Use(auth.ExtractUserID())
	//custom middleware example
	r.Use(middlewares.VerifySomething())

	if config.EnableCORS {
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

	if config.DevMod {
		r.Mount("/debug", middleware.Profiler())
	}

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	if config.ApiRequestTimeout > 0 {
		t := time.Duration(config.ApiRequestTimeout)
		r.Use(middleware.Timeout(t * time.Second))
	}

	r.Route("/dev", func(r chi.Router) {

		r.Route("/buzz", func(r chi.Router) {
			r.Post("/", handler.AddBuzz)
			r.Get("/", handler.AllBuzz)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handler.GetBuzz)
				r.Put("/", handler.ChangeBuzz)
				r.Delete("/", handler.DeleteBuzz)
			})

		})

		r.Get("/middleware", handler.GetValueFromMiddleware)

	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.ResultNotFound.Render(w)
	})

	return r
}
