package config

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/lancer-kit/armory/api"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/armory/natsx"
	"github.com/lancer-kit/uwe"
)

// Cfg main structure of the app configuration.
type Cfg struct {
	Api     api.Config   `json:"api" yaml:"api"`
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

func (cfg Cfg) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.DB, validation.Required),
		validation.Field(&cfg.ServicesInitTimeout, validation.Required),
		//uncomment this if you want to use CouchDB
		//validation.Field(&cfg.CouchDB, validation.Required),
		validation.Field(&cfg.Api, validation.Required),
		validation.Field(&cfg.NATS, validation.Required),
		validation.Field(&cfg.Workers, &uwe.WorkerExistRule{
			AvailableWorkers: AvailableWorkers,
		}),
	)
}

func (cfg Cfg) FillDefaultWorkers() {
	for k := range AvailableWorkers {
		cfg.Workers = append(cfg.Workers, k)
	}
}

type DBCfg struct {
	ConnURL     string `json:"conn_url" yaml:"conn_url"` //The database connection string.
	InitTimeout int    `json:"dbInitTimeout" yaml:"init_timeout"`
	// AutoMigrate if `true` execute db migrate up on start.
	AutoMigrate bool `json:"auto_migrate" yaml:"auto_migrate"`
	WaitForDB   bool `json:"wait_for_db" yaml:"wait_for_db"`
}

func (cfg DBCfg) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.ConnURL, validation.Required),
		validation.Field(&cfg.InitTimeout, validation.Required),
	)
}
