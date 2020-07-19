package workers

import (
	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"

	"lancer-kit/service-scaffold/config"
	"lancer-kit/service-scaffold/workers/api"
	"lancer-kit/service-scaffold/workers/foobar"
)

func InitChief(logger *logrus.Entry, cfg *config.Cfg) uwe.Chief {
	logger = logger.WithField("app_layer", "workers")

	chief := uwe.NewChief()
	chief.UseDefaultRecover()
	chief.EnableServiceSocket(config.AppInfo())
	chief.SetEventHandler(uwe.LogrusEventHandler(logger))

	chief.AddWorker(config.WorkerAPIServer,
		api.GetServer(cfg, logger.WithField("worker", config.WorkerFooBar)))

	chief.AddWorker(config.WorkerFooBar,
		foobar.NewWorker(config.WorkerFooBar, logger.WithField("worker", config.WorkerFooBar)))

	return chief
}
