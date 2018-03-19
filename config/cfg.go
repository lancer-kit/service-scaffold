package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"errors"
)

// Cfg main structure of the app configuration.
type Cfg struct {
	// DB is a database connection string.
	DB       string `yaml:"db"`
	LogLevel string `yaml:"log_level"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`

	// AutoMigrate if `true` execute db migrate up on start.
	AutoMigrate bool `yaml:"auto_migrate"`
	DevMode     bool `yaml:"dev_mode"`

	// Links are the addresses of other services
	// with which the interaction takes place.
	Links Links `yaml:"links"`

	// Services is a list of workers
	// that must be started, start all if empty.
	Services []string `yaml:"services"`
}

// Links are the addresses of other services
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

func (cfg Cfg) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.DB, validation.Required),
		validation.Field(&cfg.Host, validation.Required),
		validation.Field(&cfg.Port, validation.Required),
		validation.Field(&cfg.Links, validation.Required),
		validation.Field(&cfg.Services, serviceExistRule),
	)
}

func (cfg Cfg) FillDefaultServices() {
	for k := range AvailableServices {
		cfg.Services = append(cfg.Services, k)
	}
}

var serviceExistRule *ServiceExistRule = new(ServiceExistRule)

type ServiceExistRule struct {
	message string
}

// Validate checks if the given value is valid or not.
func (r *ServiceExistRule) Validate(value interface{}) error {
	if arr, ok := value.([]string); ok {
		for _, v := range (arr) {
			if _, ok := AvailableServices[v]; ok == false {
				return errors.New("invalid service name " + v)
			}
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
