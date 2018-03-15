package routines

import (
	"context"
	"sync"

	"github.com/inn4sc/vcg-go-common/log"
	"github.com/sirupsen/logrus"
)

// Chief is a head of workers, it must be used to register, initialize
// and correctly start and stop asynchronous executors of the type `Workman`.
type Chief struct {
	ctx         context.Context
	cancel      context.CancelFunc
	logger      *logrus.Entry
	initialized bool

	enabledWorkers map[string]struct{}
	pool           map[string]Workman
}

// AddWorkman register a new `Workman` to the `Chief` worker pool.
func (chief *Chief) AddWorkman(name string, workman Workman) {
	if chief.pool == nil {
		chief.pool = make(map[string]Workman)
	}

	chief.pool[name] = workman
}

// EnableWorkman set the worker by name enabled.
func (chief *Chief) EnableWorkman(name string) {
	if chief.enabledWorkers == nil {
		chief.enabledWorkers = make(map[string]struct{})
	}

	chief.enabledWorkers[name] = struct{}{}
}

// IsEnabled checks is enable worker with passed `name`.
func (chief *Chief) IsEnabled(name string) bool {
	if chief.enabledWorkers == nil {
		return true
	}

	_, ok := chief.enabledWorkers[name]
	return ok
}

// InitWorkers initializes all registered workers.
func (chief *Chief) InitWorkers(logger *logrus.Entry) {
	if logger == nil {
		logger = log.Default
	}

	chief.ctx, chief.cancel = context.WithCancel(context.Background())
	chief.logger = logger.WithField("service", "workman-chief")

	for name, worker := range chief.pool {
		chief.pool[name] = worker.New(chief.ctx)
	}

	chief.initialized = true
}

// Start runs all registered workers, locks until the `parentCtx` closes,
// and then gracefully stops all workers.
func (chief *Chief) Start(parentCtx context.Context) {
	if !chief.initialized {
		log.Default.Error("Workers is not initialized! Unable to start.")
		return
	}

	wg := sync.WaitGroup{}
	for name, workman := range chief.pool {
		if !chief.IsEnabled(name) {
			chief.logger.WithField("worker", name).
				Debug("Worker disabled")
			continue
		}

		wg.Add(1)
		go func(name string, workman Workman) {
			defer wg.Done()
			chief.logger.WithField("worker", name).Info("Starting worker")
			workman.Run()
		}(name, workman)
	}

	<-parentCtx.Done()
	chief.logger.Info("Begin graceful shutdown of workers")
	chief.cancel()

	wg.Wait()
	chief.logger.Info("Workers stopped")
}
