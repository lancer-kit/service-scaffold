package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/auth"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/uwe/v2/presets/api"
	"github.com/sirupsen/logrus"

	"lancer-kit/service-scaffold/config"
	"lancer-kit/service-scaffold/info"
	"lancer-kit/service-scaffold/workers/api/handler"
	"lancer-kit/service-scaffold/workers/api/middlewares"
)

func GetServer(cfg *config.Cfg, logger *logrus.Entry) *api.Server {
	return api.NewServer(cfg.Api, getRouter(logger, cfg.Api))
}

func getRouter(logger *logrus.Entry, config api.Config) http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(log.NewRequestLogger(logger.Logger))

	if config.EnableCORS {
		r.Use(getCORS().Handler)
	}

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	if config.ApiRequestTimeout > 0 {
		t := time.Duration(config.ApiRequestTimeout)
		r.Use(middleware.Timeout(t * time.Second))
	}

	r.Route("/dev", func(r chi.Router) {
		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			render.Success(w, info.App)
		})
		r.Route("/", func(r chi.Router) {
			r.Use(auth.ExtractUserID())

			r.Route("/{mId}/buzz", func(r chi.Router) {
				//custom middleware example
				r.Use(middlewares.VerifySomething())
				r.Post("/", handler.AddBuzz)
				r.Get("/", handler.AllBuzz)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", handler.GetBuzz)
					r.Put("/", handler.ChangeBuzz)
					r.Delete("/", handler.DeleteBuzz)
				})

			})
		})

		r.Route("/couch", func(r chi.Router) {
			r.Post("/", handler.AddDocument)
			r.Get("/", handler.GetAllDocument)

			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handler.GetDocument)
				r.Put("/", handler.ChangeDocument)
				r.Delete("/", handler.DeleteDocument)
			})
		})

	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.ResultNotFound.Render(w)
	})

	return r
}
func getCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "jwt", "X-UID"},
		ExposedHeaders:   []string{"Link", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
