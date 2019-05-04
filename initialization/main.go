package initialization

import (
	"fmt"
	"sync"
	"time"

	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/armory/tools"
	"github.com/lancer-kit/service-scaffold/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const flagConfig = "config"
const defaultInitInterval = 5 * time.Second

var initConfigs = map[initModule]func(*config.Cfg, *logrus.Entry) error{
	DB:   initDatabase,
	NATS: initNATS,
}

func Init(c *cli.Context) *config.Cfg {
	config.Init(c.GlobalString(flagConfig))
	cfg := config.Config()

	wg := sync.WaitGroup{}
	for i, j := range initConfigs {
		var timeout time.Duration
		if i == DB {
			timeout = time.Duration(cfg.DBInitTimeout) * time.Second
		} else {
			timeout = time.Duration(cfg.ServicesInitTimeout) * time.Second
		}

		wg.Add(1)
		go func(i initModule, j func(*config.Cfg, *logrus.Entry) error, timeout time.Duration) {
			defer wg.Done()
			ok := tools.RetryIncrementallyUntil(
				defaultInitInterval,
				timeout,

				func() bool {
					err := j(cfg, log.Default)
					if err != nil {
						log.Default.WithError(err).Error("Can't init " + i)
					}
					return err == nil
				})
			if !ok {
				log.Default.Fatal("Can't init " + i)
			}
		}(i, j, timeout)
	}
	wg.Wait()

	if cfg.AutoMigrate {
		//dbschema.SetAssets()
		count, err := db.Migrate(config.Config().DB, "up")
		if err != nil {
			log.Default.WithError(err).Error("Migrations failed")
		}
		log.Default.Info(fmt.Sprintf("Applied %d %s migration", count, "up"))
	}
	return cfg
}
