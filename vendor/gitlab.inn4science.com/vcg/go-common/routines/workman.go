package routines

import (
	"context"
	"fmt"
	"time"
)

// Worker is an interface for async workers
// which launches and manages by the `Chief`.
type Worker interface {
	// Init initializes new instance of the `Worker` implementation.
	Init(context.Context) Worker
	// Run starts the `Worker` instance execution.
	Run()
}

// DummyWorker is a simple realization of the Worker interface.
type DummyWorker struct {
	tickDuration time.Duration
	ctx          context.Context
}

// Init returns new instance of the `DummyWorker`.
func (*DummyWorker) Init(parentCtx context.Context) Worker {
	return &DummyWorker{
		ctx:          parentCtx,
		tickDuration: time.Second,
	}
}

// Run start job execution.
func (s *DummyWorker) Run() {
	ticker := time.NewTicker(15 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("I'm alive")
		case <-s.ctx.Done():
			ticker.Stop()
			fmt.Println("End job")
			return
		}
	}
}
