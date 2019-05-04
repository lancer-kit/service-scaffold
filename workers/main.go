package workers

import (
	"github.com/lancer-kit/armory/routines"
	"github.com/lancer-kit/service-scaffold/config"
	"github.com/lancer-kit/service-scaffold/workers/api"
	"github.com/lancer-kit/service-scaffold/workers/foobar"
)

var WorkerChief routines.Chief

func GetChief() *routines.Chief {
	WorkerChief = routines.Chief{EnableByDefault: true}
	WorkerChief.AddWorker(config.WorkerAPIServer, api.Server())
	WorkerChief.AddWorker(config.WorkerFooBar, &foobar.Service{})

	return &WorkerChief
}
