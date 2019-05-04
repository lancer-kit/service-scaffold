package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/lancer-kit/armory/api"
	"github.com/lancer-kit/armory/api/infoworker"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/armory/natsx"
)

// Cfg main structure of the app configuration.
type Cfg struct {
	DB                  string           `json:"db" yaml:"db"` // DB is a database connection string.
	DBInitTimeout       int              `json:"dbInitTimeout" yaml:"db_init_timeout"`
	ServicesInitTimeout int              `json:"servicesInitTimeout" yaml:"services_init_timeout"`
	CouchDB             string           `json:"couchdb" yaml:"couchdb"` // CouchDB is a couchdb url connection string.
	Api                 api.Config       `json:"api" yaml:"api"`
	InfoWorker          *infoworker.Conf `yaml:"info_worker"`

	// AutoMigrate if `true` execute db migrate up on start.
	AutoMigrate bool `json:"auto_migrate" yaml:"auto_migrate"`
	DevMode     bool `json:"dev_mode" yaml:"dev_mode"`
	WaitForDB   bool `json:"wait_for_db" yaml:"wait_for_db"`

	NATS natsx.Config `json:"nats" yaml:"nats"`
	Log  log.Config   `json:"log" yaml:"log"`

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
		validation.Field(&cfg.DBInitTimeout, validation.Required),
		validation.Field(&cfg.ServicesInitTimeout, validation.Required),
		//uncomment this if you want to use CouchDB
		//validation.Field(&cfg.CouchDB, validation.Required),
		validation.Field(&cfg.Api, validation.Required),
		validation.Field(&cfg.Links, validation.Required),
		validation.Field(&cfg.NATS, validation.Required),
		validation.Field(&cfg.Workers, &WorkerExistRule{
			AvailableWorkers: AvailableWorkers,
		}),
	)
}

func (cfg Cfg) FillDefaultWorkers() {
	for k := range AvailableWorkers {
		cfg.Workers = append(cfg.Workers, k)
	}
}
