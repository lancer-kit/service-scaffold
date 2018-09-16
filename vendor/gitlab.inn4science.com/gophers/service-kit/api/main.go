package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/sirupsen/logrus"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-kit/routines"
)

var ForceStopTimeout = 5 * time.Second

// Config is a parameters for `http.Server`.
type Config struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`

	ApiRequestTimeout int  `json:"api_request_timeout" yaml:"api_request_timeout"`
	DevMod            bool `json:"dev_mod" yaml:"dev_mod"`
	EnableCORS        bool `json:"enable_cors" yaml:"enable_cors"`

	RestartOnFail bool `json:"restart_on_fail" yaml:"restart_on_fail"`
}

//Validate - Validate config required fields
func (c *Config) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Host, validation.Required),
		validation.Field(&c.Port, validation.Required),
	)
}

// TCPAddr returns tcp address for server.
func (c *Config) TCPAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// Server
type Server struct {
	Name      string
	Config    Config
	GetConfig func() Config
	GetRouter func(*logrus.Entry, Config) http.Handler

	ctx    context.Context
	logger *logrus.Entry
}

func NewServer(name string, config Config, rGetter func(*logrus.Entry, Config) http.Handler) Server {
	return Server{
		Name:      name,
		Config:    config,
		GetRouter: rGetter,
	}
}

func (s *Server) Init(parentCtx context.Context) routines.Worker {
	var ok bool
	s.logger, ok = parentCtx.Value(routines.CtxKeyLog).(*logrus.Entry)
	if !ok {
		s.logger = log.Default
	}

	if s.Name == "" {
		s.Name = "api-server"
	}

	s.logger = s.logger.WithField("service", s.Name)
	s.ctx = parentCtx
	return s
}

func (s *Server) RestartOnFail() bool {
	if s.GetConfig != nil {
		return s.GetConfig().RestartOnFail
	}
	return s.Config.RestartOnFail
}

func (s *Server) Run() {
	if s.GetConfig != nil {
		s.Config = s.GetConfig()
	}
	addr := fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: s.GetRouter(s.logger, s.Config),
	}

	go func() {
		s.logger.Info("Starting API Server at: ", addr)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.WithError(err).Error("server failed")
		}
	}()

	<-s.ctx.Done()
	s.logger.Info("Shutting down the API Server...")
	serverCtx, _ := context.WithTimeout(context.Background(), ForceStopTimeout)
	server.Shutdown(serverCtx)
	s.logger.Info("Api Server gracefully stopped")
}
