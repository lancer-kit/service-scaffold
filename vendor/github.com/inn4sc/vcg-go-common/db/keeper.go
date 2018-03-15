package db

import (
	"context"
	"time"

	"github.com/inn4sc/vcg-go-common/routines"
)

// Keeper is a service that verifies the connection
// to the database using `ping` at a specified interval.
type Keeper struct {
	// Conn is a database connection interface.
	Conn *SQLConn
	// Interval is a time duration between ping calls.
	Interval time.Duration
	// Name is a service name
	Name string

	ctx context.Context
}

// New inits worker instance.
func (s *Keeper) New(parentCtx context.Context) routines.Workman {
	if s.Conn == nil {
		s.Conn = conn
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
