package main

import (
	"fmt"

	"github.com/urfave/cli"
	"gitlab.inn4science.com/vcg/go-common/log"

	"gitlab.inn4science.com/vcg/go-skeleton/config"
	"gitlab.inn4science.com/vcg/go-skeleton/dbschema"
)

func migrateDB(cfgPath string, direction dbschema.MigrateDir) *cli.ExitError {
	initConfig(cfgPath)
	count, err := dbschema.Migrate(config.Config().DB, direction)
	if err != nil {
		log.Default.WithError(err).Error("Migrations failed")
		return cli.NewExitError(fmt.Sprintf("migration %s failed", direction), 1)
	}

	log.Default.Info(fmt.Sprintf("Applied %d %s migration", count, direction))
	return nil
}

var migrateCommand = cli.Command{
	Name:  "migrate",
	Usage: "apply db migration",

	Subcommands: []cli.Command{
		{
			Name:  "up",
			Usage: "apply up migration direction",
			Flags: cfgFlag,
			Action: func(c *cli.Context) error {
				err := migrateDB(c.String("config"), dbschema.MigrateUp)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "down",
			Usage: "drop and clean database schema",
			Flags: cfgFlag,
			Action: func(c *cli.Context) error {
				err := migrateDB(c.String("config"), dbschema.MigrateDown)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "redo",
			Usage: "reset database schema",
			Flags: cfgFlag,
			Action: func(c *cli.Context) error {
				err := migrateDB(c.String("config"), dbschema.MigrateDown)
				if err != nil {
					return err
				}
				err = migrateDB(c.String("config"), dbschema.MigrateUp)
				if err != nil {
					return err
				}

				return nil
			},
		},
	},
}
