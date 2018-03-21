package workers

import (
	"gitlab.inn4science.com/vcg/go-common/routines"
	"github.com/inn4sc/go-skeleton/config"
	"github.com/inn4sc/go-skeleton/workers/api"
	"github.com/inn4sc/go-skeleton/workers/foobar"
)

var WorkerChief routines.Chief

func init() {
	WorkerChief = routines.Chief{}
	WorkerChief.AddWorker(config.WorkerAPIServer, &api.Server{})
	WorkerChief.AddWorker(config.WorkerFooBar, &foobar.Service{})
}
