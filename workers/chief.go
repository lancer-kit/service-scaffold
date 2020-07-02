package workers

import (
	"github.com/lancer-kit/uwe/v2"
	"github.com/sirupsen/logrus"

	"lancer-kit/service-scaffold/config"
	"lancer-kit/service-scaffold/workers/api"
	"lancer-kit/service-scaffold/workers/foobar"
)

func InitChief(logger *logrus.Entry, cfg *config.Cfg) uwe.Chief {
	chief := uwe.NewChief()
	chief.UseDefaultRecover()
	chief.SetEventHandler(func(event uwe.Event) {
		var level logrus.Level
		switch event.Level {
		case uwe.LvlFatal, uwe.LvlError:
			level = logrus.ErrorLevel
		case uwe.LvlInfo:
			level = logrus.InfoLevel
		default:
			level = logrus.WarnLevel
		}

		logger.WithFields(event.Fields).
			Log(level, event.Message)
	})

	logger = logger.WithField("app_layer", "workers")

	chief.AddWorker(config.WorkerAPIServer, api.GetServer(cfg, logger.WithField("worker", config.WorkerFooBar)))
	chief.AddWorker(config.WorkerFooBar, foobar.NewWorker(config.WorkerFooBar, logger.WithField("worker", config.WorkerFooBar)))

	return chief
}
