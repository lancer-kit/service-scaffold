package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"gitlab.inn4science.com/gophers/service-kit/tools"
)

// Links are the addresses of other micro services
// with which the interaction takes place.
type Links struct {
	UserAPI      tools.URL `yaml:"user_api"`
	PaymentGate  tools.URL `yaml:"payment_gate"`
	Rate         tools.URL `yaml:"rate"`
	Transactions tools.URL `yaml:"transactions"`
}

func (links Links) Validate() error {
	return validation.ValidateStruct(&links,
		validation.Field(&links.Rate, validation.Required),
		validation.Field(&links.PaymentGate, validation.Required),
		validation.Field(&links.UserAPI, validation.Required),
		validation.Field(&links.Transactions, validation.Required),
	)
}
