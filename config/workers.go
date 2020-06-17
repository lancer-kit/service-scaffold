package config

import (
	"errors"
	"reflect"

	"github.com/lancer-kit/uwe/v2"
)

const (
	WorkerInfoServer uwe.WorkerName = "info-server"
	WorkerAPIServer  uwe.WorkerName = "api-server"
	WorkerDBKeeper   uwe.WorkerName = "db-keeper"
	WorkerFooBar     uwe.WorkerName = "foobar"
)

var AvailableWorkers = map[uwe.WorkerName]struct{}{
	WorkerInfoServer: {},
	WorkerDBKeeper:   {},
	WorkerAPIServer:  {},
	WorkerFooBar:     {},
}

type WorkerExistRule struct {
	message          string
	AvailableWorkers map[uwe.WorkerName]struct{}
}

// Validate checks that service exist on the system
func (r *WorkerExistRule) Validate(value interface{}) error {
	if value == nil || reflect.ValueOf(value).IsNil() {
		return nil
	}
	arr, ok := value.([]uwe.WorkerName)
	if !ok {
		return errors.New("can't convert list of workers to []string")
	}

	for _, v := range arr {
		if _, ok := r.AvailableWorkers[uwe.WorkerName(v)]; !ok {
			return errors.New("invalid service name " + string(v))
		}
	}

	return nil
}

// Error sets the error message for the rule.
func (r *WorkerExistRule) Error(message string) *WorkerExistRule {
	return &WorkerExistRule{
		message: message,
	}
}
