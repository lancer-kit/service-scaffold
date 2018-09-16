package config

import (
	"errors"
)

const (
	WorkerInfoServer = "info-server"
	WorkerAPIServer  = "api-server"
	WorkerDBKeeper   = "db-keeper"
	WorkerFooBar     = "foobar"
)

var AvailableWorkers = map[string]struct{}{
	WorkerInfoServer: {},
	WorkerDBKeeper:   {},
	WorkerAPIServer:  {},
	WorkerFooBar:     {},
}

type WorkerExistRule struct {
	message          string
	AvailableWorkers map[string]struct{}
}

// Validate checks that service exist on the system
func (r *WorkerExistRule) Validate(value interface{}) error {
	arr, ok := value.([]string)
	if !ok {
		return errors.New("can't convert list of workers to []string")
	}
	for _, v := range arr {
		if _, ok := r.AvailableWorkers[v]; !ok {
			return errors.New("invalid service name " + v)
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
