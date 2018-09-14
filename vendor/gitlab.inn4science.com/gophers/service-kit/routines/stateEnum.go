package routines

type WorkerState int32

const (
	WorkerWrongStateChange WorkerState = -1
	WorkerNull             WorkerState = iota
	WorkerDisabled
	WorkerPresent
	WorkerEnabled
	WorkerInitialized
	WorkerRun
	WorkerStopped
	WorkerFailed
)
