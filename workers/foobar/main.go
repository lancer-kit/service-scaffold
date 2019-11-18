package foobar

import (
	"context"
	"fmt"
	"time"

	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/uwe"
	"github.com/sirupsen/logrus"
)

type Worker struct {
	name   string
	logger *logrus.Entry
}

func (d Worker) Init(ctx context.Context) uwe.Worker {
	var rawValue = ctx.Value(uwe.CtxKeyLog)

	logger, ok := rawValue.(*logrus.Entry)
	if !ok {
		logger = log.Default
	}

	return &Worker{
		name:   d.name,
		logger: logger,
	}
}

func (Worker) RestartOnFail() bool {
	return true
}

func (d Worker) Run(wCtx uwe.WContext) uwe.ExitCode {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			d.logger.Info("Perform my task")
		case m := <-wCtx.MessageBus():
			d.logger.
				WithField("Sender", m.Sender).
				WithField("Target", m.Target).
				WithField("data", fmt.Sprintf("%+v", m.Data)).
				Info("got new message")
		case <-wCtx.Done():
			ticker.Stop()
			d.logger.Info("Receive exit code, stop working")
			return uwe.ExitCodeOk
		}
	}

}
