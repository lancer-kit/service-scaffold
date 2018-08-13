package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"gitlab.inn4science.com/gophers/service-kit/log"

	"gitlab.inn4science.com/gophers/service-kit/db"
	"gitlab.inn4science.com/gophers/service-scaffold/config"
	"gitlab.inn4science.com/gophers/service-scaffold/dbschema"
)

var migrateCommand = cli.Command{
	Name:  "migrate",
	Usage: "apply db migration",

	Subcommands: []cli.Command{
		{
			Name:  "up",
			Usage: "apply up migration direction",
			Flags: cfgFlag,
			Action: func(c *cli.Context) error {
				err := migrateDB(c.String("config"), db.MigrateUp)
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
				err := migrateDB(c.String("config"), db.MigrateDown)
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
				err := migrateDB(c.String("config"), db.MigrateDown)
				if err != nil {
					return err
				}
				err = migrateDB(c.String("config"), db.MigrateUp)
				if err != nil {
					return err
				}

				return nil
			},
		},
	},
}

func migrateDB(cfgPath string, direction db.MigrateDir) *cli.ExitError {
	config.Init(cfgPath)

	dbschema.SetAssets()

	count, err := db.Migrate(config.Config().DB, direction)
	if err != nil {
		log.Default.WithError(err).Error("Migrations failed")
		return cli.NewExitError(fmt.Sprintf("migration %s failed", direction), 1)
	}

	log.Default.Info(fmt.Sprintf("Applied %d %s migration", count, direction))
	return nil
}
