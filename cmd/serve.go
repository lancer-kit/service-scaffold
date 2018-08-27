package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"gitlab.inn4science.com/gophers/service-kit/db"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
	"gitlab.inn4science.com/gophers/service-scaffold/workers"
)

var serveCommand = cli.Command{
	Name:   "serve",
	Usage:  "starts " + config.ServiceName + " workers",
	Flags:  cfgFlag,
	Action: serveAction,
}

func serveAction(c *cli.Context) error {
	config.Init(c.String("config"))
	cfg := config.Config()

	if cfg.AutoMigrate {
		count, err := db.Migrate(config.Config().DB, "up")
		if err != nil {
			log.Default.WithError(err).Error("Migrations failed")
		}
		log.Default.Info(fmt.Sprintf("Applied %d %s migration", count, "up"))
	}

	workers.GetChief().RunAll(cfg.Log.AppName, cfg.Workers...)
	return nil
}
