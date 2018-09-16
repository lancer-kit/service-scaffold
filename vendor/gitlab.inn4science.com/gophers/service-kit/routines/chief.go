package routines

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.inn4science.com/gophers/service-kit/log"
)

// CtxKey is the type of context keys for the values placed by`Chief`.
type CtxKey string

const (
	// CtxKeyLog is a context key for a `*logrus.Entry` value.
	CtxKeyLog CtxKey = "chief-log"
)

// ForceStopTimeout is a timeout for killing all workers.
var ForceStopTimeout = 45 * time.Second

// Chief is a head of workers, it must be used to register, initialize
// and correctly start and stop asynchronous executors of the type `Worker`.
type Chief struct {
	ctx    context.Context
	cancel context.CancelFunc
	logger *logrus.Entry
	active bool
	wg     sync.WaitGroup
	wPool  WorkerPool
	// EnableByDefault sets all the working `Enabled`
	// if none of the workers is passed on to enable.
	EnableByDefault bool
}

// AddWorker register a new `Worker` to the `Chief` worker pool.
func (chief *Chief) AddWorker(name string, worker Worker) {
	chief.wPool.SetWorker(name, worker)
}

// EnableWorkers enables all worker from the `names` list.
// By default, all added workers are enabled. After the first call
// of this method, only directly enabled workers will be active
func (chief *Chief) EnableWorkers(names ...string) {
	for _, name := range names {
		chief.wPool.EnableWorker(name)
	}

	if len(names) == 0 && chief.EnableByDefault {
		for name := range chief.wPool.workers {
			chief.wPool.EnableWorker(name)
		}
	}
}

// EnableWorker enables the worker with the specified `name`.
// By default, all added workers are enabled. After the first call
// of this method, only directly enabled workers will be active
func (chief *Chief) EnableWorker(name string) {
	chief.wPool.EnableWorker(name)
}

// IsEnabled checks is enable worker with passed `name`.
func (chief *Chief) IsEnabled(name string) bool {
	return chief.wPool.IsEnabled(name)
}

// InitWorkers initializes all registered workers.
func (chief *Chief) InitWorkers(logger *logrus.Entry) {
	if logger == nil {
		logger = log.Default
	}

	chief.logger = logger.WithField("service", "worker-chief")
	chief.ctx, chief.cancel = context.WithCancel(context.Background())
	chief.ctx = context.WithValue(chief.ctx, CtxKeyLog, chief.logger)

	for name := range chief.wPool.states {
		chief.wPool.InitWorker(name, chief.ctx)
	}

	chief.active = true
}

// Start runs all registered workers, locks until the `parentCtx` closes,
// and then gracefully stops all workers.
func (chief *Chief) Start(parentCtx context.Context) {
	if !chief.active {
		log.Default.Error("Workers is not initialized! Unable to start.")
		return
	}

	chief.wg = sync.WaitGroup{}
	for name, worker := range chief.wPool.workers {
		if !chief.wPool.IsEnabled(name) {
			chief.logger.WithField("worker", name).
				Debug("Worker disabled")
			continue
		}

		chief.wg.Add(1)
		go chief.runWorker(name, worker)
	}

	<-parentCtx.Done()
	chief.logger.Info("Begin graceful shutdown of workers")
	chief.active = false
	chief.cancel()

	chief.wg.Wait()
	chief.logger.Info("Workers stopped")
}

// RunAll start worker pool and lock context
// until it intercepts `syscall.SIGTERM`, `syscall.SIGINT`.
// NOTE: Use this method ONLY as a top-level action.
func (chief *Chief) RunAll(appName string, workers ...string) error {
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	chief.EnableWorkers(workers...)

	chief.InitWorkers(log.Default)
	go func() {
		defer close(done)
		chief.Start(ctx)
	}()

	chief.logger.Info(appName + " started")

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	exitSignal := <-gracefulStop
	chief.logger.WithField("signal", exitSignal).
		Info("Received signal. Terminating service...")

	cancel()

	select {
	case <-done:
		chief.logger.Info("Graceful exit.")
		return nil
	case <-time.NewTimer(ForceStopTimeout).C:
		chief.logger.Warn("Graceful exit timeout exceeded. Force exit.")
		return errors.New("Graceful exit timeout exceeded")
	}
}

func (chief *Chief) runWorker(name string, worker Worker) {
	defer chief.wg.Done()

	defer func() {
		err := recover()
		if err == nil {
			return
		}
	}()

startWorker:
	chief.logger.WithField("worker", name).Info("Starting worker")

	err := chief.wPool.RunWorkerExec(name)
	if err != nil {
		chief.logger.WithField("worker", name).
			WithError(err).
			Error("Worker failed")

		if worker.RestartOnFail() && chief.active {
			time.Sleep(time.Second)
			goto startWorker
		}
	}

	chief.wPool.StopWorker(name)
}

func (chief *Chief) GetWorkersStates() map[string]WorkerState {
	return chief.wPool.GetWorkersStates()
}

func (chief *Chief) GetContext() context.Context {
	return chief.ctx
}
