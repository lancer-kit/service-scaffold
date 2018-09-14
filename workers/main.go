package workers

import (
	"context"

	"gitlab.inn4science.com/gophers/service-kit/api/infoworker"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-kit/routines"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
	"gitlab.inn4science.com/gophers/service-scaffold/info"
	"gitlab.inn4science.com/gophers/service-scaffold/workers/api"
	"gitlab.inn4science.com/gophers/service-scaffold/workers/foobar"
)

var WorkerChief routines.Chief

func GetChief() *routines.Chief {
	WorkerChief = routines.Chief{EnableByDefault: true}

	//create context with pointer to cihef
	ctx := context.Background()
	ctx = context.WithValue(ctx, "chief", &WorkerChief)

	worker := infoworker.GetInfoWorker(
		config.Config().InfoWorker,
		ctx,
		info.App,
	)

	//add worker only in dev mode
	if config.Config().Api.DevMod == true {
		log.Default.Info("Starting info worker")
		WorkerChief.AddWorker(config.WorkerInfoServer, worker)
		if !WorkerChief.EnableByDefault {
			WorkerChief.EnableWorker(config.WorkerInfoServer)
		}

	} else {
		log.Default.Info("Info worker can be started only in dev mode")
	}
	WorkerChief.AddWorker(config.WorkerAPIServer, api.Server())
	WorkerChief.AddWorker(config.WorkerFooBar, &foobar.Service{})
	return &WorkerChief
}
