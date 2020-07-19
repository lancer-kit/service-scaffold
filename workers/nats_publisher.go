package workers

import (
	"time"

	"lancer-kit/service-scaffold/models"

	"github.com/lancer-kit/armory/natsx"
	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"
)

type NATSPub struct {
	name   uwe.WorkerName
	logger *logrus.Entry
	bus    <-chan models.Event
}

func NewWorker(logger *logrus.Entry, bus <-chan models.Event) *NATSPub {
	return &NATSPub{
		logger: logger,
		bus:    bus,
	}
}

func (d *NATSPub) Init() error {
	return nil
}

func (d *NATSPub) Run(wCtx uwe.Context) error {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case msg := <-d.bus:
			d.logger.Info("got new event")

			err := natsx.PublishMessage(models.NATSTopic, msg)
			if err != nil {
				d.logger.WithError(err).Error("unable to publish event")
			}

		case <-wCtx.Done():
			ticker.Stop()
			d.logger.Info("Receive exit code, stop working")
			return nil
		}
	}

}
