package routines

import (
	"context"
	"fmt"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gitlab.inn4science.com/gophers/service-kit/log"
)

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

// RestartOnFail determines the need to restart the worker, if it stopped
func (s *DummyWorker) RestartOnFail() bool {
	return true
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
func TestChief_InitWorkers(t *testing.T) {
	require := require.New(t)

	wp := WorkerPool{}
	worker := new(DummyWorker)
	wp.SetWorker("dummyWorker", worker)

	testChief := new(Chief)
	testChief.wPool = wp

	tests := []struct {
		name  string
		chief *Chief
	}{
		{
			name:  "Test init workers",
			chief: testChief,
		},
	}
	for _, tt := range tests {
		log.Default.Info(fmt.Sprintf("Started %s", tt.name))

		tt.chief.InitWorkers(nil)
		require.NotNilf(tt.chief.logger, "Chief.logger is not initialized")
		require.NotNilf(tt.chief.ctx, "Chief.ctx is not initialized")
		require.NotNilf(tt.chief.cancel, "Chief.cansel is not initialized")
		require.Truef(tt.chief.active, "Chief not active")

		log.Default.Info(fmt.Sprintf("%s finished successfully", tt.name))
	}
}

func TestChief_Start(t *testing.T) {
	require := require.New(t)

	wp := WorkerPool{}
	worker := new(DummyWorker)
	wp.SetWorker("dummyWorker", worker)

	testChief := new(Chief)
	testChief.wPool = wp

	ctx, cansel := context.WithCancel(context.Background())

	tests := []struct {
		name       string
		chief      *Chief
		parentCtx  context.Context
		canselFunc context.CancelFunc
	}{
		{
			name:       "Test start chief",
			chief:      testChief,
			parentCtx:  ctx,
			canselFunc: cansel,
		},
	}

	for _, tt := range tests {

		log.Default.Info(fmt.Sprintf("Started %s", tt.name))

		tt.chief.InitWorkers(nil)

		go tt.chief.Start(tt.parentCtx)

		require.NotNilf(tt.chief.wg, "Error chief is not active")

		time.Sleep(20 * time.Second)
		tt.canselFunc()
		time.Sleep(5 * time.Second)

		require.Falsef(tt.chief.active, "Chief is still active after shuttdowning of workers")
		log.Default.Info(fmt.Sprintf("%s finished successfully", tt.name))
	}
}

func TestChief_RunAll(t *testing.T) {
	require := require.New(t)

	workerPool := WorkerPool{}
	worker := new(DummyWorker)
	workerPool.SetWorker("dummyWorker", worker)

	testChief := new(Chief)
	testChief.wPool = workerPool

	ctx, cansel := context.WithCancel(context.Background())

	tests := []struct {
		name       string
		chief      *Chief
		parentCtx  context.Context
		canselFunc context.CancelFunc
	}{
		{
			name:       "Test run chief",
			chief:      testChief,
			parentCtx:  ctx,
			canselFunc: cansel,
		},
	}

	for _, tt := range tests {

		log.Default.Info(fmt.Sprintf("Started %s", tt.name))

		go func() {
			time.Sleep(20 * time.Second)
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		}()

		err := tt.chief.RunAll("Test chief", "dummyWorker")
		require.NoError(err)
		require.Falsef(tt.chief.active, "Chief is still active after shuttdowning of workers")
		log.Default.Info(fmt.Sprintf("%s finished successfully", tt.name))
	}
}

//restart indicates if worker is need to be restarted or not
var restart bool

type RestartWorker struct {
	tickDuration time.Duration
	ctx          context.Context
}

// Init returns new instance of the `RestartWorker`.
func (*RestartWorker) Init(parentCtx context.Context) Worker {
	return &RestartWorker{
		ctx:          parentCtx,
		tickDuration: time.Second,
	}
}

// RestartOnFail determines the need to restart the worker, if it stopped
func (s *RestartWorker) RestartOnFail() bool {
	return restart
}

// Run start job execution.
func (s *RestartWorker) Run() {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			//create panic for tests
			panic("planned panic when executing worker")
			fmt.Println("I'm alive")
		case <-s.ctx.Done():
			ticker.Stop()
			fmt.Println("End job")
			return
		}
	}
}

func TestRestartOnFailWorker(t *testing.T) {
	require := require.New(t)

	wp := WorkerPool{}
	worker := new(RestartWorker)
	wp.SetWorker("restartWorker", worker)

	testChief := new(Chief)
	testChief.wPool = wp

	ctxRest, canselRest := context.WithCancel(context.Background())
	ctxNoRest, canselNoRest := context.WithCancel(context.Background())

	tests := []struct {
		name       string
		chief      *Chief
		parentCtx  context.Context
		canselFunc context.CancelFunc
		ifRestart  bool
	}{
		{
			name:       "Run with restart on fail",
			chief:      testChief,
			parentCtx:  ctxRest,
			canselFunc: canselRest,
			ifRestart:  true,
		},
		{
			name:       "Run with no restart on fail",
			chief:      testChief,
			parentCtx:  ctxNoRest,
			canselFunc: canselNoRest,
			ifRestart:  false,
		},
	}

	for _, tt := range tests {

		log.Default.Info(fmt.Sprintf("Started %s", tt.name))

		tt.chief.InitWorkers(nil)
		restart = tt.ifRestart

		go tt.chief.Start(tt.parentCtx)

		require.NotNilf(tt.chief.wg, "Error chief is not active")

		time.Sleep(30 * time.Second)
		tt.canselFunc()
		time.Sleep(5 * time.Second)

		require.Falsef(tt.chief.active, "Chief is still active after shuttdowning of workers")
		log.Default.Info(fmt.Sprintf("%s finished successfully", tt.name))
	}
}
