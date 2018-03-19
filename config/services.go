package config

const (
	ServiceAPIServer = "api-server"
	ServiceDBKeeper  = "db-keeper"
	ServiceFooBar    = "foobar"
)

var AvailableServices = map[string]struct{}{
	ServiceDBKeeper:  {},
	ServiceAPIServer: {},
	ServiceFooBar:    {},
}
