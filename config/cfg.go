package config

import (
	"io/ioutil"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/armory/natsx"
	"github.com/lancer-kit/uwe/v2"
	"github.com/lancer-kit/uwe/v2/presets/api"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const ServiceName = "service-scaffold"

// Cfg main structure of the app configuration.
type Cfg struct {
	API     api.Config   `json:"api" yaml:"api"`
	DB      DBCfg        `json:"db" yaml:"db"`           // DB is a database connection string.
	CouchDB string       `json:"couchdb" yaml:"couchdb"` // CouchDB is a couchdb url connection string.
	NATS    natsx.Config `json:"nats" yaml:"nats"`
	Log     log.Config   `json:"log" yaml:"log"`

	DevMode             bool `json:"dev_mode" yaml:"dev_mode"`
	ServicesInitTimeout int  `json:"servicesInitTimeout" yaml:"services_init_timeout"`

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
		validation.Field(&cfg.Workers, &WorkerExistRule{
			AvailableWorkers: GetAvailableWorkers(),
		}),
	)
}

func (cfg Cfg) FillDefaultWorkers() {
	for k := range GetAvailableWorkers() {
		cfg.Workers = append(cfg.Workers, k)
	}
}

type DBCfg struct {
	ConnURL     string `json:"conn_url" yaml:"conn_url"` // The database connection string.
	InitTimeout int    `json:"dbInitTimeout" yaml:"init_timeout"`
	// AutoMigrate if `true` execute db migrate up on start.
	AutoMigrate bool `json:"auto_migrate" yaml:"auto_migrate"`
	WaitForDB   bool `json:"wait_for_db" yaml:"wait_for_db"`
}

// Validate is an implementation of Validatable interface from ozzo-validation.
func (cfg DBCfg) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.ConnURL, validation.Required),
		validation.Field(&cfg.InitTimeout, validation.Required),
	)
}

func ReadConfig(path string) Cfg {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.New().WithError(err).
			WithField("path", path).
			Fatal("unable to read config file")
	}

	config := new(Cfg)
	err = yaml.Unmarshal(rawConfig, config)
	if err != nil {
		logrus.New().WithError(err).
			WithField("raw_config", rawConfig).
			Fatal("unable to unmarshal config file")
	}

	err = config.Validate()
	if err != nil {
		logrus.New().WithError(err).
			Fatal("Invalid configuration")
	}

	_, err = log.Init(log.Config{
		AppName:  config.Log.AppName,
		Level:    config.Log.Level,
		Sentry:   config.Log.Sentry,
		AddTrace: config.Log.AddTrace,
		JSON:     config.Log.JSON,
	})
	if err != nil {
		logrus.New().
			WithError(err).
			Fatal("Unable to init log")
	}
	return *config
}
