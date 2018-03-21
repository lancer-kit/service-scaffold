package config

const (
	WorkerAPIServer = "api-server"
	WorkerDBKeeper  = "db-keeper"
	WorkerFooBar    = "foobar"
)

var AvailableWorkers = map[string]struct{}{
	WorkerDBKeeper:  {},
	WorkerAPIServer: {},
	WorkerFooBar:    {},
}
