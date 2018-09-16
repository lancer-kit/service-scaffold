package routines

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

var (
	ErrWorkerNotInitialized = errors.New("worker not initialized")
)

// WorkerPool is
type WorkerPool struct {
	workers map[string]Worker
	states  map[string]WorkerState
	rw      sync.RWMutex
}

// DisableWorker sets state `WorkerDisabled` for workers with the specified `name`.
func (pool *WorkerPool) DisableWorker(name string) {
	pool.SetState(name, WorkerDisabled)
}

// EnableWorker sets state `WorkerEnabled` for workers with the specified `name`.
func (pool *WorkerPool) EnableWorker(name string) {
	if s := pool.states[name]; s != WorkerPresent {
		pool.SetState(name, WorkerWrongStateChange)
		return
	}
	pool.SetState(name, WorkerEnabled)
}

// StartWorker sets state `WorkerEnabled` for workers with the specified `name`.
func (pool *WorkerPool) StartWorker(name string) {
	if s := pool.states[name]; s != WorkerStopped &&
		s != WorkerInitialized && s != WorkerFailed {
		pool.SetState(name, WorkerWrongStateChange)
		return
	}
	pool.SetState(name, WorkerRun)
}

// StopWorker sets state `WorkerStopped` for workers with the specified `name`.
func (pool *WorkerPool) StopWorker(name string) {
	if s := pool.states[name]; s != WorkerRun && s != WorkerFailed {
		pool.SetState(name, WorkerWrongStateChange)
		return
	}
	pool.SetState(name, WorkerStopped)
}

// FailWorker sets state `WorkerFailed` for workers with the specified `name`.
func (pool *WorkerPool) FailWorker(name string) {
	pool.SetState(name, WorkerFailed)
}

// IsEnabled checks is enable worker with passed `name`.
func (pool *WorkerPool) IsEnabled(name string) bool {
	if pool.states == nil {
		return false
	}

	state := pool.states[name]
	return state >= WorkerEnabled
}

// IsAlive checks is active worker with passed `name`.
func (pool *WorkerPool) IsAlive(name string) bool {
	if pool.states == nil {
		return false
	}

	state := pool.states[name]
	return state == WorkerRun
}

// InitWorker initializes all present workers.
func (pool *WorkerPool) InitWorker(name string, ctx context.Context) {
	pool.rw.Lock()
	defer pool.rw.Unlock()
	pool.check()

	if pool.states[name] < WorkerEnabled {
		return
	}

	pool.workers[name] = pool.workers[name].Init(ctx)
	pool.states[name] = WorkerInitialized
}

// SetState updates state of specified worker.
func (pool *WorkerPool) SetState(name string, state WorkerState) {
	pool.rw.Lock()
	defer pool.rw.Unlock()

	pool.check()
	pool.states[name] = state
}

// SetWorker adds worker into pool.
func (pool *WorkerPool) SetWorker(name string, worker Worker) {
	pool.rw.Lock()
	defer pool.rw.Unlock()

	pool.check()
	pool.workers[name] = worker
	pool.states[name] = WorkerPresent
}

// RunWorkerExec adds worker into pool.
func (pool *WorkerPool) RunWorkerExec(name string) (err error) {
	defer func() {
		rec := recover()
		if rec == nil {
			return
		}
		pool.FailWorker(name)

		var ok bool
		err, ok = rec.(error)
		if !ok {
			err = fmt.Errorf("%v", rec)
		}
	}()

	if s := pool.states[name]; s != WorkerInitialized {
		return ErrWorkerNotInitialized
	}

	pool.StartWorker(name)
	pool.workers[name].Run()
	pool.StopWorker(name)

	return
}

func (pool *WorkerPool) check() {
	if pool.states == nil {
		pool.states = make(map[string]WorkerState)
	}
	if pool.workers == nil {
		pool.workers = make(map[string]Worker)
	}
}

func (pool *WorkerPool) GetWorkersStates() map[string]WorkerState {
	return pool.states
}
