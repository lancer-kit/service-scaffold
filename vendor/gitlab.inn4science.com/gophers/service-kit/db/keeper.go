package db

import (
	"context"
	"time"

	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-kit/routines"
)

// Keeper is a service that verifies the connection
// to the database using `ping` at a specified interval.
type Keeper struct {
	// сonn is a database connection interface.
	Conn *SQLConn
	// Interval is a time duration between ping calls.
	Interval time.Duration
	// Name is a service name
	Name string

	ctx context.Context
}

// Init inits worker instance.
func (s *Keeper) Init(parentCtx context.Context) routines.Worker {
	if s.Conn == nil {
		s.Conn = GetConnector()
	}

	if s.Conn.logger == nil {
		s.Conn.logger = log.Get()
	}

	if s.Interval == 0 {
		s.Interval = 5 * time.Second
	}

	if s.Name == "" {
		s.Name = "db-keeper"
	}

	s.ctx = parentCtx
	return s
}

// RestartOnFail determines the need to restart the worker, if it stopped.
func (s *Keeper) RestartOnFail() bool {
	return true
}

// Run starts worker execution.
func (s *Keeper) Run() {
	ticker := time.NewTicker(s.Interval)
	logger := s.Conn.logger.WithField("service", s.Name)
	var err error

	for {
		select {
		case <-ticker.C:
			logger.Debug("ping db conn")
			if err = s.Conn.db.Ping(); err != nil {
				logger.WithError(err).Error("")
			}
		case <-s.ctx.Done():
			return
		}
	}
}
