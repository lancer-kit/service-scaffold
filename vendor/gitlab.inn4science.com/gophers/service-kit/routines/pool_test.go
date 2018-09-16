package routines

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.inn4science.com/gophers/service-kit/log"
)

type test struct {
	name       string
	pool       *WorkerPool
	workerName string
	want       WorkerState
}

//wrapper for WorkerPool methods
type wrapper func(wp *WorkerPool, name string)

// mockWorker is an implementation
type mockWorker struct {
}

func (m mockWorker) Init(context.Context) Worker {
	return m
}
func (m mockWorker) RestartOnFail() bool {
	return false
}
func (m mockWorker) Run() {

}

// initPool returns WorkerPool instance suitable for most tests
func initPool() *WorkerPool {
	newPool := new(WorkerPool)
	newPool.workers = make(map[string]Worker)
	newPool.workers["test"] = mockWorker{}

	return newPool
}

//runTestCases runs all test cases in particular test func
func runTestCases(t *testing.T, testCases []test, w wrapper) {
	for _, tt := range testCases {
		log.Default.Info(fmt.Sprintf("Started %s", tt.name))

		w(tt.pool, tt.workerName)
		require.Equalf(t, tt.want, tt.pool.states[tt.workerName], fmt.Sprintf("Error in: %s", tt.name))

		log.Default.Info(fmt.Sprintf("%s finished successfully", tt.name))
	}
}

func TestWorkerPool_DisableWorker(t *testing.T) {

	testPool := initPool()
	tests := []test{
		{
			name:       "Disable worker test",
			pool:       testPool,
			workerName: "test",
			want:       WorkerDisabled,
		},
	}
	runTestCases(t, tests, func(wp *WorkerPool, name string) {
		wp.DisableWorker(name)
	})

}

func TestWorkerPool_EnableWorker(t *testing.T) {
	testPool := initPool()

	tests := []test{
		{
			name:       "Enable worker test",
			pool:       testPool,
			workerName: "test",
			want:       WorkerEnabled,
		},
	}
	runTestCases(t, tests, func(wp *WorkerPool, name string) {
		wp.EnableWorker(name)
	})
}

func TestWorkerPool_StartWorker(t *testing.T) {
	testPool := initPool()

	tests := []test{
		{
			name:       "Start worker test",
			pool:       testPool,
			workerName: "test",
			want:       WorkerRun,
		},
	}

	runTestCases(t, tests, func(wp *WorkerPool, name string) {
		wp.StartWorker(name)
	})
}

func TestWorkerPool_StopWorker(t *testing.T) {
	testPool := initPool()

	tests := []test{
		{
			name:       "Stop worker test",
			pool:       testPool,
			workerName: "test",
			want:       WorkerStopped,
		},
	}

	runTestCases(t, tests, func(wp *WorkerPool, name string) {
		wp.StopWorker(name)
	})
}

func TestWorkerPool_FailWorker(t *testing.T) {
	testPool := initPool()

	tests := []test{
		{
			name:       "Fail worker test",
			pool:       testPool,
			workerName: "test",
			want:       WorkerFailed,
		},
	}

	runTestCases(t, tests, func(wp *WorkerPool, name string) {
		wp.FailWorker(name)
	})
}

func TestWorkerPool_IsEnabled(t *testing.T) {
	testPoolNil := initPool()

	testPoolNotExist := initPool()
	testPoolNotExist.states = make(map[string]WorkerState)

	testPoolOk := initPool()
	testPoolOk.states = make(map[string]WorkerState)
	testPoolOk.states["test"] = WorkerEnabled

	tests := []struct {
		name       string
		pool       *WorkerPool
		workerName string
		want       bool
	}{
		{
			name:       "Test with nil states",
			pool:       testPoolNil,
			workerName: "test",
			want:       false,
		},
		{
			name:       "Test with non existing worker",
			pool:       testPoolNotExist,
			workerName: "notExists",
			want:       false,
		},
		{
			name:       "Test with correct data",
			pool:       testPoolOk,
			workerName: "test",
			want:       true,
		},
	}

	for _, tt := range tests {
		log.Default.Info(fmt.Sprintf("Started %s", tt.name))

		res := tt.pool.IsEnabled(tt.workerName)
		require.Equalf(t, tt.want, res, fmt.Sprintf("Error in: %s", tt.name))

		log.Default.Info(fmt.Sprintf("%s finished successfully", tt.name))
	}
}

func TestWorkerPool_IsAlive(t *testing.T) {
	testPoolNil := initPool()

	testPoolNotExist := initPool()
	testPoolNotExist.states = make(map[string]WorkerState)

	testPoolOk := initPool()
	testPoolOk.states = make(map[string]WorkerState)
	testPoolOk.states["test"] = WorkerRun

	tests := []struct {
		name       string
		pool       *WorkerPool
		workerName string
		want       bool
	}{
		{
			name:       "Test with nil states",
			pool:       testPoolNil,
			workerName: "test",
			want:       false,
		},
		{
			name:       "Test with non existing worker",
			pool:       testPoolNotExist,
			workerName: "notExists",
			want:       false,
		},
		{
			name:       "Test with correct data",
			pool:       testPoolOk,
			workerName: "test",
			want:       true,
		},
	}

	for _, tt := range tests {
		log.Default.Info(fmt.Sprintf("Started %s", tt.name))

		res := tt.pool.IsAlive(tt.workerName)
		require.Equalf(t, tt.want, res, fmt.Sprintf("Error in: %s", tt.name))

		log.Default.Info(fmt.Sprintf("%s finished successfully", tt.name))
	}

}

func TestWorkerPool_InitWorker(t *testing.T) {
	testPoolDisabled := initPool()

	testPoolPresent := initPool()
	testPoolPresent.states = make(map[string]WorkerState)
	testPoolPresent.states["test"] = WorkerPresent

	tests := []test{
		{
			name:       "Test with status disabled",
			pool:       testPoolDisabled,
			workerName: "test",
			want:       WorkerDisabled,
		},
		{
			name:       "Test with status present",
			pool:       testPoolPresent,
			workerName: "test",
			want:       WorkerInitialized,
		},
	}

	runTestCases(t, tests, func(wp *WorkerPool, name string) {
		wp.InitWorker(name, nil)
	})
}

func TestWorkerPool_SetState(t *testing.T) {
	testPool := initPool()
	tests := []test{
		{
			name:       `Set "disabled" state test`,
			pool:       testPool,
			workerName: "test",
			want:       WorkerDisabled,
		},
		{
			name:       `Set "enabled" state test`,
			pool:       testPool,
			workerName: "test",
			want:       WorkerEnabled,
		},
	}

	for _, tt := range tests {
		log.Default.Info(fmt.Sprintf("Started %s", tt.name))

		tt.pool.SetState(tt.workerName, tt.want)
		require.Equalf(t, tt.want, tt.pool.states[tt.workerName], fmt.Sprintf("Error in: %s", tt.name))

		log.Default.Info(fmt.Sprintf("%s finished successfully", tt.name))
	}

}

func TestWorkerPool_SetWorker(t *testing.T) {

	testPool := new(WorkerPool)

	type args struct {
		name   string
		worker Worker
	}

	tests := []test{
		{
			name:       "Set worker test",
			pool:       testPool,
			workerName: "test",
			want:       WorkerPresent,
		},
	}

	runTestCases(t, tests, func(wp *WorkerPool, name string) {
		wp.SetWorker(name, mockWorker{})
	})

}

func TestWorkerPool_RunWorkerExec(t *testing.T) {
	testPool := initPool()
	tests := []test{
		{
			name:       `Set "disabled" state test`,
			pool:       testPool,
			workerName: "test",
			want:       WorkerStopped,
		},
	}

	for _, tt := range tests {
		log.Default.Info(fmt.Sprintf("Started %s", tt.name))

		err := tt.pool.RunWorkerExec(tt.workerName)

		require.NoError(t, err, fmt.Sprintf("Error in: %s", tt.name))
		require.Equalf(t, tt.want, tt.pool.states[tt.workerName], fmt.Sprintf("Error in: %s", tt.name))

		log.Default.Info(fmt.Sprintf("%s finished successfully", tt.name))
	}
}
