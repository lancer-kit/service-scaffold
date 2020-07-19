package config

import (
	"github.com/lancer-kit/uwe/v2"
)

const (
	WorkerAPIServer     uwe.WorkerName = "api-server"
	WorkerNATSPublisher uwe.WorkerName = "nats-publisher"
)

func availableWorkers() map[uwe.WorkerName]struct{} {
	return map[uwe.WorkerName]struct{}{
		WorkerAPIServer:     {},
		WorkerNATSPublisher: {},
	}
}
