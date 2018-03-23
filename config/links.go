package config

import "github.com/go-ozzo/ozzo-validation"

// Links are the addresses of other micro services
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

