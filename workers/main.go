package workers

import (
	"context"

	"gitlab.inn4science.com/gophers/service-kit/api/infoworker"
	"gitlab.inn4science.com/gophers/service-kit/routines"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
	"gitlab.inn4science.com/gophers/service-scaffold/info"
	"gitlab.inn4science.com/gophers/service-scaffold/workers/api"
	"gitlab.inn4science.com/gophers/service-scaffold/workers/foobar"
)

var WorkerChief routines.Chief

func GetChief() *routines.Chief {
	WorkerChief = routines.Chief{EnableByDefault: true}
	WorkerChief.AddWorker(config.WorkerAPIServer, api.Server())
	WorkerChief.AddWorker(config.WorkerFooBar, &foobar.Service{})

	//add worker only in dev mode
	if config.Config().InfoWorker != nil {
		//create context with pointer to chief
		ctx := context.Background()
		ctx = context.WithValue(ctx, "chief", &WorkerChief)
		worker := infoworker.GetInfoWorker(*config.Config().InfoWorker, ctx, info.App)
		WorkerChief.AddWorker(config.WorkerInfoServer, worker)
		if !WorkerChief.EnableByDefault {
			WorkerChief.EnableWorker(config.WorkerInfoServer)
		}

	}
	return &WorkerChief
}
