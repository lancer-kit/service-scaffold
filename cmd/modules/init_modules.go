package modules

import (
	"fmt"
	"time"

	"lancer-kit/service-scaffold/config"
	"lancer-kit/service-scaffold/repo/schema"

	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/initialization"
	"github.com/lancer-kit/armory/natsx"
	cdb "github.com/leesper/couchdb-golang"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Init(c *cli.Context) (*config.Cfg, error) {
	cfg, err := config.ReadConfig(c.GlobalString(config.FlagConfig))
	if err != nil {
		return nil, errors.Wrap(err, "unable to read config")
	}

	err = getModules(cfg).InitAll()
	if err != nil {
		return nil, errors.Wrap(err, "modules initialization failed")
	}

	return &cfg, nil
}

func getModules(cfg config.Cfg) initialization.Modules {
	return initialization.Modules{
		initialization.Module{
			Name:         "postgres",
			DependsOn:    "",
			Timeout:      time.Duration(cfg.DB.InitTimeout) * time.Second,
			InitInterval: 500 * time.Millisecond,
			Init: func(entry *logrus.Entry) error {
				err := db.Init(cfg.DB.ConnectionString(), entry)
				if err != nil {
					return errors.Wrap(err, "db init failed")
				}

				if cfg.DB.AutoMigrate {
					count, err := schema.Migrate(cfg.DB.ConnectionString(), "up")
					if err != nil {
						return errors.Wrap(err, "auto-migration failed")
					}
					entry.Info(fmt.Sprintf("Applied %d %s migration", count, "up"))
				}

				return nil
			},
		},

		initialization.Module{
			Name:         "nats",
			DependsOn:    "",
			Timeout:      time.Duration(cfg.ServicesInitTimeout) * time.Second,
			InitInterval: 500 * time.Millisecond,
			Init: func(entry *logrus.Entry) error {
				natsx.SetConfig(&cfg.NATS)

				_, err := natsx.GetConn()
				if err != nil {
					return errors.Wrap(err, "nats init failed")
				}

				return nil
			},
		},

		initialization.Module{
			Name:         "couchdb",
			DependsOn:    "",
			Timeout:      time.Duration(cfg.ServicesInitTimeout) * time.Second,
			InitInterval: 500 * time.Millisecond,
			Init: func(entry *logrus.Entry) error {
				_, err := cdb.NewDatabase(cfg.CouchDB)
				if err != nil {
					return errors.Wrap(err, "unable to connect to couchdb")
				}
				return nil
			},
		},
	}

}
