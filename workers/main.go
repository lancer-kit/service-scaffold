package workers

import (
	"gitlab.inn4science.com/gophers/service-kit/routines"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
	"gitlab.inn4science.com/gophers/service-scaffold/workers/api"
	"gitlab.inn4science.com/gophers/service-scaffold/workers/foobar"
)

var WorkerChief routines.Chief

func init() {
	WorkerChief = routines.Chief{}
	WorkerChief.AddWorker(config.WorkerAPIServer, api.Server())
	WorkerChief.AddWorker(config.WorkerFooBar, &foobar.Service{})
}
