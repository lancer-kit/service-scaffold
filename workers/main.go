package workers

import (
	"gitlab.inn4science.com/vcg/go-common/routines"
	"gitlab.inn4science.com/vcg/go-skeleton/config"
	"gitlab.inn4science.com/vcg/go-skeleton/workers/api"
	"gitlab.inn4science.com/vcg/go-skeleton/workers/foobar"
)

var WorkerChief routines.Chief

func init() {
	WorkerChief = routines.Chief{}
	WorkerChief.AddWorker(config.WorkerAPIServer, &api.Server{})
	WorkerChief.AddWorker(config.WorkerFooBar, &foobar.Service{})
}
