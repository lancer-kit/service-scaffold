package natswrap

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
)

// Config is configuration for the interaction with the NATS server.
type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (config Config) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.Host, validation.Required),
		validation.Field(&config.Port, validation.Required),
	)
}

func (config Config) ToURL() string {
	return fmt.Sprintf("nats://%s:%d", config.Host, config.Port)
}
