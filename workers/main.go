package workers

import (
	"gitlab.inn4science.com/internal/service-scaffold/config"
	"gitlab.inn4science.com/internal/service-scaffold/workers/api"
	"gitlab.inn4science.com/internal/service-scaffold/workers/foobar"
	"gitlab.inn4science.com/vcg/go-common/routines"
)

var WorkerChief routines.Chief

func init() {
	WorkerChief = routines.Chief{}
	WorkerChief.AddWorker(config.WorkerAPIServer, &api.Server{})
	WorkerChief.AddWorker(config.WorkerFooBar, &foobar.Service{})
}
