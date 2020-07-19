package cmd

import (
	"fmt"

	"lancer-kit/service-scaffold/config"
	"lancer-kit/service-scaffold/dbschema"

	"github.com/lancer-kit/armory/db"
	"github.com/lancer-kit/armory/log"
	"github.com/urfave/cli"
)

func migrateCmd() cli.Command {
	var migrateCommand = cli.Command{
		Name:  "migrate",
		Usage: "apply db migration",

		Subcommands: []cli.Command{
			{
				Name:  "up",
				Usage: "apply up migration direction",
				Action: func(c *cli.Context) error {
					cfg, err := config.ReadConfig(c.GlobalString(config.FlagConfig))
					if err != nil {
						return cli.NewExitError(err, 1)
					}

					return migrateDB(db.MigrateUp, cfg)
				},
			},
			{
				Name:  "down",
				Usage: "drop and clean database schema",
				Action: func(c *cli.Context) error {
					cfg, err := config.ReadConfig(c.GlobalString(config.FlagConfig))
					if err != nil {
						return cli.NewExitError(err, 1)
					}

					return migrateDB(db.MigrateDown, cfg)
				},
			},
			{
				Name:  "redo",
				Usage: "reset database schema",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "force, f",
						Usage: "if the flag set, the database schema will be dropped instead of migration Down",
					},
				},
				Action: func(c *cli.Context) error {
					cfg, err := config.ReadConfig(c.GlobalString(config.FlagConfig))
					if err != nil {
						return cli.NewExitError(err, 1)
					}

					if c.Bool("force") {
						err = dbschema.DropSchema(cfg.DB.ConnectionString())
					} else {
						err = migrateDB(db.MigrateDown, cfg)
					}
					if err != nil {
						return cli.NewExitError(err, 1)
					}

					return migrateDB(db.MigrateUp, cfg)
				},
			},
		},
	}
	return migrateCommand
}

func migrateDB(direction db.MigrateDir, cfg config.Cfg) *cli.ExitError {
	count, err := dbschema.Migrate(cfg.DB.ConnectionString(), direction)
	if err != nil {
		log.Get().WithError(err).Error("Migrations failed")
		return cli.NewExitError(fmt.Sprintf("migration %s failed", direction), 1)
	}

	log.Get().Info(fmt.Sprintf("Applied %d %s migration", count, direction))
	return nil
}
