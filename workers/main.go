package workers

import (
	"gitlab.inn4science.com/gophers/service-kit/routines"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
	"gitlab.inn4science.com/gophers/service-scaffold/workers/api"
	"gitlab.inn4science.com/gophers/service-scaffold/workers/foobar"
)

func GetChief() *routines.Chief {
	chief := &routines.Chief{}
	chief.AddWorker(config.WorkerAPIServer, api.Server())
	chief.AddWorker(config.WorkerFooBar, &foobar.Service{})
	return chief
}
