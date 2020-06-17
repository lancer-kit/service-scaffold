package foobar

import (
	"time"

	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"
)

type Worker struct {
	name   uwe.WorkerName
	logger *logrus.Entry
}

func NewWorker(name uwe.WorkerName, logger *logrus.Entry) *Worker {
	return &Worker{
		name:   name,
		logger: logger,
	}
}

func (d Worker) Init() error {
	return nil
}

func (d Worker) Run(wCtx uwe.Context) error {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			d.logger.Info("Perform my task")
		case <-wCtx.Done():
			ticker.Stop()
			d.logger.Info("Receive exit code, stop working")
			return nil
		}
	}

}
