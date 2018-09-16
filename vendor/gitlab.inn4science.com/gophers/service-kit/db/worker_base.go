package db

import (
	"context"

	"github.com/sirupsen/logrus"
)

// WBase is a base structure for worker which
// uses database and need transactions support.
type WBase struct {
	DB     Transactional
	Logger *logrus.Entry
	Ctx    context.Context
	Err    error
}

func (s *WBase) DBTxRollback() {
	if s.Err == nil {
		return
	}

	if !s.DB.IsInTx() {
		return
	}

	s.Logger.WithError(s.Err).Info("job failed, try to rollback db")
	rbErr := s.DB.Rollback()
	if rbErr != nil {
		s.Logger.WithError(rbErr).Fatal("failed to rollback db")
	}
}

func (s *WBase) Recover() {
	err := recover()
	if err == nil {
		return
	}
	s.Logger.WithField("panic", err).Error("Caught panic")

	if s.DB.IsInTx() {
		s.DBTxRollback()
	}
}
