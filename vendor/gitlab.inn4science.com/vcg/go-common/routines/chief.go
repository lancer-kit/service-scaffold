package routines

import (
	"context"
	"sync"

	"gitlab.inn4science.com/vcg/go-common/log"
	"github.com/sirupsen/logrus"
)

// CtxKey is the type of context keys for the values placed by`Chief`.
type CtxKey string

const (
	// CtxKeyLog is a context key for a `*logrus.Entry` value.
	CtxKeyLog CtxKey = "chief-log"
)

// Chief is a head of workers, it must be used to register, initialize
// and correctly start and stop asynchronous executors of the type `Worker`.
type Chief struct {
	ctx         context.Context
	cancel      context.CancelFunc
	logger      *logrus.Entry
	initialized bool

	enabledWorkers map[string]struct{}
	pool           map[string]Worker
}

// AddWorker register a new `Worker` to the `Chief` worker pool.
func (chief *Chief) AddWorker(name string, worker Worker) {
	if chief.pool == nil {
		chief.pool = make(map[string]Worker)
	}

	chief.pool[name] = worker
}

// EnableWorkers enables all worker from the `names` list.
// By default, all added workers are enabled. After the first call
// of this method, only directly enabled workers will be active
func (chief *Chief) EnableWorkers(names ...string) {
	if chief.enabledWorkers == nil {
		chief.enabledWorkers = make(map[string]struct{})
	}

	for _, name := range names {
		chief.enabledWorkers[name] = struct{}{}
	}
}

// EnableWorker enables the worker with the specified `name`.
// By default, all added workers are enabled. After the first call
// of this method, only directly enabled workers will be active
func (chief *Chief) EnableWorker(name string) {
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

	chief.logger = logger.WithField("service", "worker-chief")
	chief.ctx, chief.cancel = context.WithCancel(context.Background())
	chief.ctx = context.WithValue(chief.ctx, CtxKeyLog, chief.logger)

	for name, worker := range chief.pool {
		chief.pool[name] = worker.Init(chief.ctx)
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
	for name, worker := range chief.pool {
		if !chief.IsEnabled(name) {
			chief.logger.WithField("worker", name).
				Debug("Worker disabled")
			continue
		}

		wg.Add(1)
		go func(name string, worker Worker) {
			defer wg.Done()
			chief.logger.WithField("worker", name).Info("Starting worker")
			worker.Run()
		}(name, worker)
	}

	<-parentCtx.Done()
	chief.logger.Info("Begin graceful shutdown of workers")
	chief.cancel()

	wg.Wait()
	chief.logger.Info("Workers stopped")
}
