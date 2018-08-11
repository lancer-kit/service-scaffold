package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-kit/natswrap"
)

// Cfg main structure of the app configuration.
type Cfg struct {
	// DB is a database connection string.
	DB       string          `yaml:"db"`
	LogLevel string          `yaml:"log_level"`
	Host     string          `yaml:"host"`
	Port     int             `yaml:"port"`
	NATS     natswrap.Config `yaml:"nats"`

	Log log.Config `yaml:"log"`

	// AutoMigrate if `true` execute db migrate up on start.
	AutoMigrate       bool `yaml:"auto_migrate"`
	DevMode           bool `yaml:"dev_mode"`
	WaitForDB         bool `yaml:"wait_for_db"`
	EnableCORS        bool `yaml:"enable_cors"`
	ApiRequestTimeout int  `yaml:"api_request_timeout"`

	// Links are the addresses of other services
	// with which the interaction takes place.
	Links Links `yaml:"links"`

	// Workers is a list of workers
	// that must be started, start all if empty.
	Workers []string `yaml:"workers"`
}

func (cfg Cfg) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.DB, validation.Required),
		validation.Field(&cfg.Host, validation.Required),
		validation.Field(&cfg.Port, validation.Required),
		validation.Field(&cfg.Links, validation.Required),
		validation.Field(&cfg.NATS, validation.Required),
		validation.Field(&cfg.Workers, new(WorkerExistRule)),
	)
}

func (cfg Cfg) FillDefaultWorkers() {
	for k := range AvailableWorkers {
		cfg.Workers = append(cfg.Workers, k)
	}
}
