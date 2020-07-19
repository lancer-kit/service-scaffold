package workers

import (
	"lancer-kit/service-scaffold/models"

	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"

	"lancer-kit/service-scaffold/config"
	"lancer-kit/service-scaffold/workers/api"
)

func InitChief(logger *logrus.Entry, cfg *config.Cfg) uwe.Chief {
	logger = logger.WithField("app_layer", "workers")

	chief := uwe.NewChief()
	chief.UseDefaultRecover()
	chief.EnableServiceSocket(config.AppInfo())
	chief.SetEventHandler(uwe.LogrusEventHandler(logger))

	eventBus := make(chan models.Event, 16)
	chief.AddWorker(config.WorkerAPIServer,
		api.GetServer(cfg, logger.WithField("worker", config.WorkerAPIServer), eventBus))

	chief.AddWorker(config.WorkerNATSPublisher,
		NewWorker(logger.WithField("worker", config.WorkerNATSPublisher), eventBus))

	return chief
}
