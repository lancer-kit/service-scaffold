package infoworker

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.inn4science.com/gophers/service-kit/api"
	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-kit/routines"
)

//structure with info about service
type Info struct {
	App     string `json:"app"`
	Version string `json:"version"`
	Tag     string `json:"tag"`
	Build   string `json:"build"`
}

//structure with configuration for worker
type Conf struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Profiler bool   `json:"profiler" yaml:"profiler"`
	Prefix   string `json:"prefix" yaml:"prefix"`
}

type InfoWorker struct {
	api.Server

	logger    *logrus.Entry
	parentCtx context.Context // context with pointer to chief which started this worker

	Info   Info
	Config Conf
}

//GetInfoWorker returns initialized info worker struct. !Context must contain pointer to worker chief!
func GetInfoWorker(cfg Conf, ctx context.Context, info Info) *InfoWorker {
	res := &InfoWorker{
		parentCtx: ctx,
		Info:      info,
		Config:    cfg,
		Server: api.Server{
			Name: "info-server",
			Config: api.Config{
				Host:              cfg.Host,
				Port:              cfg.Port,
				ApiRequestTimeout: 60,
				EnableCORS:        false,
				RestartOnFail:     true,
			},
		},
	}

	res.GetRouter = res.GetInfoRouter
	return res
}

//GetInfoRouter returns workers api
func (iw *InfoWorker) GetInfoRouter(logger *logrus.Entry, cfg api.Config) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Timeout(time.Duration(cfg.ApiRequestTimeout) * time.Second))
	r.Use(middleware.Recoverer)
	r.Use(log.NewRequestLogger(logger.Logger))

	prefix := "/"
	if iw.Config.Prefix != "" {
		prefix = iw.Config.Prefix
	}

	r.Route(prefix, func(r chi.Router) {
		if iw.Config.Profiler {
			r.Mount("/debug", middleware.Profiler())
		}

		r.Route("/info", func(r chi.Router) {
			r.Get("/", iw.Version)
			r.Get("/workers", iw.Workers)
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.ResultNotFound.Render(w)
	})

	return r
}

//Version is a handler which responses with Info structure
func (iw *InfoWorker) Version(w http.ResponseWriter, r *http.Request) {
	if iw.Info == (Info{}) {
		err := errors.New("Info must not be empty!")
		iw.logger.Error(err)
		render.ResultServerError.SetError(err).Render(w)
		return
	}

	render.Success(w, iw.Info)
}

//Workers is a handler which responses with JSON with all workers in parent chief
func (iw *InfoWorker) Workers(w http.ResponseWriter, r *http.Request) {
	parentChief := iw.parentCtx.Value("chief").(*routines.Chief)
	if parentChief == nil {
		err := errors.New("Context must not be empty!")
		iw.logger.Error(err)
		render.ResultServerError.SetError(err).Render(w)
		return
	}

	workers := parentChief.GetWorkersStates()
	if len(workers) == 0 {
		iw.logger.Debug("No workers are currently running")
	}

	render.Success(w, workers)
}
