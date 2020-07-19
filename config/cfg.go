package config

import (
	"io/ioutil"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/armory/natsx"
	"github.com/lancer-kit/uwe/v2"
	"github.com/lancer-kit/uwe/v2/presets/api"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const FlagConfig = "config"

// Cfg main structure of the app configuration.
type Cfg struct {
	API     api.Config      `yaml:"api"`
	DB      db.SecureConfig `yaml:"db"`      // DB is a database connection string.
	CouchDB string          `yaml:"couchdb"` // CouchDB is a couchdb url connection string.
	NATS    natsx.Config    `yaml:"nats"`
	Log     log.NConfig     `yaml:"log"`

	DevMode             bool `yaml:"dev_mode"`
	ServicesInitTimeout int  `yaml:"services_init_timeout"`

	// Workers is a list of workers
	// that must be started, start all if empty.
	Workers []uwe.WorkerName `yaml:"workers"`
}

// Validate is an implementation of Validatable interface from ozzo-validation.
func (cfg Cfg) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.DB, validation.Required),
		validation.Field(&cfg.ServicesInitTimeout, validation.Required),
		validation.Field(&cfg.CouchDB, validation.Required),
		validation.Field(&cfg.API, validation.Required),
		validation.Field(&cfg.NATS, validation.Required),
		validation.Field(&cfg.Workers, &uwe.WorkerExistRule{AvailableWorkers: availableWorkers()}),
	)
}

func (cfg Cfg) FillDefaultWorkers() {
	for k := range availableWorkers() {
		cfg.Workers = append(cfg.Workers, k)
	}
}

func ReadConfig(path string) (Cfg, error) {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return Cfg{}, errors.Wrap(err, "unable to read config file")
	}

	config := new(Cfg)
	err = yaml.Unmarshal(rawConfig, config)
	if err != nil {
		return Cfg{}, errors.Wrap(err, "unable to unmarshal config file")
	}

	err = config.Validate()
	if err != nil {
		return Cfg{}, errors.Wrap(err, "invalid configuration")
	}

	_, err = log.Init(log.Config{
		AppName:  config.Log.AppName,
		Level:    config.Log.Level.Get(),
		Sentry:   config.Log.Sentry.Get(),
		AddTrace: config.Log.AddTrace,
		JSON:     config.Log.JSON,
	})
	if err != nil {
		return Cfg{}, errors.Wrap(err, "unable to init log")

	}

	return *config, nil
}
