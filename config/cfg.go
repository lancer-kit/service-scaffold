package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"gitlab.inn4science.com/vcg/go-common/log"
	"errors"
)

// Cfg main structure of the app configuration.
type Cfg struct {
	// DB is a database connection string.
	DB       string `yaml:"db"`
	LogLevel string `yaml:"log_level"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`

	Log log.Config `yaml:"log"`

	// AutoMigrate if `true` execute db migrate up on start.
	AutoMigrate bool `yaml:"auto_migrate"`
	DevMode     bool `yaml:"dev_mode"`
	WaitForDB   bool `yaml:"wait_for_db"`
	EnableCORS  bool `yaml:"enable_cors"`

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
		validation.Field(&cfg.Workers, serviceExistRule),
	)
}

func (cfg Cfg) FillDefaultServices() {
	for k := range AvailableWorkers {
		cfg.Workers = append(cfg.Workers, k)
	}
}

// Links are the addresses of other workers
// with which the interaction takes place.
type Links struct {
	UserAPI      string `yaml:"user_api"`
	PaymentGate  string `yaml:"payment_gate"`
	Rate         string `yaml:"rate"`
	Transactions string `yaml:"transactions"`
}

func (links Links) Validate() error {
	return validation.ValidateStruct(&links,
		validation.Field(&links.Rate, validation.Required),
		validation.Field(&links.PaymentGate, validation.Required),
		validation.Field(&links.UserAPI, validation.Required),
		validation.Field(&links.Transactions, validation.Required),
	)
}

type ServiceExistRule struct {
	message string
}

// Validate checks that service exist on the system
func (r *ServiceExistRule) Validate(value interface{}) error {
	arr, ok := value.([]string)
	if !ok {
		return errors.New("can't convert list of workers to []string")
	}
	for _, v := range arr {
		if _, ok := AvailableWorkers[v]; !ok {
			return errors.New("invalid service name " + v)
		}
	}
	return nil
}

// Error sets the error message for the rule.
func (r *ServiceExistRule) Error(message string) *ServiceExistRule {
	return &ServiceExistRule{
		message: message,
	}
}

var serviceExistRule = new(ServiceExistRule)
