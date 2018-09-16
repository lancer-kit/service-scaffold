package routines

import (
	"context"
)

// Worker is an interface for async workers
// which launches and manages by the `Chief`.
type Worker interface {
	// Init initializes new instance of the `Worker` implementation.
	Init(context.Context) Worker
	// RestartOnFail determines the need to restart the worker, if it stopped.
	RestartOnFail() bool
	// Run starts the `Worker` instance execution.
	Run() //todo(mike): add result or error
}
